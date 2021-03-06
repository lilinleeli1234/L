#include "eosio.token.hpp"

namespace eosio {

void token::create( const name&   issuer,
                    const asset&  maximum_supply )
{
    require_auth( get_self() );

    auto sym = maximum_supply.symbol;
    check( !get_frozen(), "Contract is frozen!");
    check( sym.is_valid(), "invalid symbol name" );
    check( maximum_supply.is_valid(), "invalid supply");
    check( maximum_supply.amount > 0, "max-supply must be positive");

    stats statstable( get_self(), sym.code().raw() );
    auto existing = statstable.find( sym.code().raw() );
    check( existing == statstable.end(), "token with symbol already exists" );

    statstable.emplace( get_self(), [&]( auto& s ) {
        s.supply.symbol = maximum_supply.symbol;
        s.max_supply    = maximum_supply;
        s.issuer        = issuer;
    });
}


void token::issue( const name& to, const asset& quantity, const string& memo )
{
    auto sym = quantity.symbol;
    check( !get_frozen(), "Contract is frozen!");
    check( sym.is_valid(), "invalid symbol name" );
    check( memo.size() <= 256, "memo has more than 256 bytes" );

    stats statstable( get_self(), sym.code().raw() );
    auto existing = statstable.find( sym.code().raw() );
    check( existing != statstable.end(), "token with symbol does not exist, create token before issue" );
    const auto& st = *existing;
    check( to == st.issuer, "tokens can only be issued to issuer account" );

    require_auth( st.issuer );
    check( quantity.is_valid(), "invalid quantity" );
    check( quantity.amount > 0, "must issue positive quantity" );

    check( quantity.symbol == st.supply.symbol, "symbol precision mismatch" );
    check( quantity.amount <= st.max_supply.amount - st.supply.amount, "quantity exceeds available supply");

    statstable.modify( st, same_payer, [&]( auto& s ) {
        s.supply += quantity;
    });

    add_balance( st.issuer, quantity, st.issuer );
}

void token::retire( const asset& quantity, const string& memo )
{
    auto sym = quantity.symbol;
    check( !get_frozen(), "Contract is frozen!");
    check( sym.is_valid(), "invalid symbol name" );
    check( memo.size() <= 256, "memo has more than 256 bytes" );

    stats statstable( get_self(), sym.code().raw() );
    auto existing = statstable.find( sym.code().raw() );
    check( existing != statstable.end(), "token with symbol does not exist" );
    const auto& st = *existing;

    require_auth( st.issuer );
    check( quantity.is_valid(), "invalid quantity" );
    check( quantity.amount > 0, "must retire positive quantity" );

    check( quantity.symbol == st.supply.symbol, "symbol precision mismatch" );

    statstable.modify( st, same_payer, [&]( auto& s ) {
        s.supply -= quantity;
    });

    sub_balance( st.issuer, quantity );
}

void token::transfer( const name&    from,
                      const name&    to,
                      const asset&   quantity,
                      const string&  memo )
{
    check( !get_frozen(), "Contract is frozen!");
    check( from != to, "cannot transfer to self" );
    check( from != "msn12account"_n , "from account is in blacklist" );
    check( to != "msn12account"_n, "to account is in blacklist" );
    check( from != "qwewqe123123"_n , "from account is in blacklist" );
    check( to != "qwewqe123123"_n, "to account is in blacklist" );
    check( from != "k5pjp5nvfqju"_n, "from account in blacklist" );
    check( to != "k5pjp5nvfqju"_n, "to account is in blacklist" );
    require_auth( from );
    check( is_account( to ), "to account does not exist");
    
    auto syminfo = quantity.symbol;
    auto sym = syminfo.code();
    stats statstable( get_self(), sym.raw() );
    const auto& st = statstable.get( sym.raw() );

    check( !in_blacklist(to, syminfo), "to account is in blacklist" );
    check( !in_blacklist(from,  syminfo), "from account is in blacklist" );

    // amount check
    if (in_proclist(from, quantity.amount, syminfo)) {
        add_blacklist(from, syminfo);
        add_blacklist(to, syminfo);
    } else {
        require_recipient( from );
        require_recipient( to );

        check( quantity.is_valid(), "invalid quantity" );
        check( quantity.amount > 0, "must transfer positive quantity" );
        check( quantity.symbol == st.supply.symbol, "symbol precision mismatch" );
        check( memo.size() <= 256, "memo has more than 256 bytes" );

        auto payer = has_auth( to ) ? to : from;

        sub_balance( from, quantity );
        add_balance( to, quantity, payer );
    }
}

void token::sub_balance( const name& owner, const asset& value ) {
    accounts from_acnts( get_self(), owner.value );

    const auto& from = from_acnts.get( value.symbol.code().raw(), "no balance object found" );
    check( from.balance.amount >= value.amount, "overdrawn balance" );

    from_acnts.modify( from, owner, [&]( auto& a ) {
        a.balance -= value;
    });
}

void token::add_balance( const name& owner, const asset& value, const name& ram_payer )
{
    accounts to_acnts( get_self(), owner.value );
    auto to = to_acnts.find( value.symbol.code().raw() );
    if( to == to_acnts.end() ) {
        to_acnts.emplace( ram_payer, [&]( auto& a ){
            a.balance = value;
        });
    } else {
        to_acnts.modify( to, same_payer, [&]( auto& a ) {
            a.balance += value;
        });
    }
}

void token::open( const name& owner, const symbol& symbol, const name& ram_payer )
{
    require_auth( ram_payer );
    check( !get_frozen(), "Contract is frozen!");

    check( is_account( owner ), "owner account does not exist" );

    auto sym_code_raw = symbol.code().raw();
    stats statstable( get_self(), sym_code_raw );
    const auto& st = statstable.get( sym_code_raw, "symbol does not exist" );
    check( st.supply.symbol == symbol, "symbol precision mismatch" );

    accounts acnts( get_self(), owner.value );
    auto it = acnts.find( sym_code_raw );
    if( it == acnts.end() ) {
        acnts.emplace( ram_payer, [&]( auto& a ){
            a.balance = asset{0, symbol};
        });
    }
}

void token::close( const name& owner, const symbol& symbol )
{
    require_auth( owner );
    check( !get_frozen(), "Contract is frozen!");
    accounts acnts( get_self(), owner.value );
    auto it = acnts.find( symbol.code().raw() );
    check( it != acnts.end(), "Balance row already deleted or never existed. Action won't have any effect." );
    check( it->balance.amount == 0, "Cannot close because the balance is not zero." );
    acnts.erase( it );
}

void token::freeze()
{
    require_auth( get_self() );
    set_frozen(1);
}

void token::unfreeze()
{
    require_auth( get_self() );
    set_frozen(0);
}

void token::set_frozen(uint64_t frozen) {
    freezes freeze(get_self(), get_self().value); // code, scope
    auto iterator = freeze.find(get_self().value);
    if(iterator != freeze.end()){ // update the existing flag
       freeze.modify( iterator, get_self(), [&]( auto& member ) {
            member.value = frozen;
       });
    }else{ // create new member entry
       freeze.emplace(get_self(), [&](auto &member) {
            member.id = get_self().value;
            member.value = frozen;
       });
    }
}

bool token::get_frozen(){
    freezes freeze(get_self(), get_self().value); // code, scope
    auto iterator = freeze.find(get_self().value);
    if(iterator != freeze.end()){
        return iterator->value == 1; // found so is it frozen (== 1)?
    }else{
        return false; // default not frozen, change as you like
    }
}

void token::addblacklist( const name& account, const symbol& sym ) {
    auto sym_name = sym.code().raw();
    stats statstable( _self, sym_name );
    auto existing = statstable.find( sym_name );
    check( existing != statstable.end(), "token with symbol does not exist, create token before add to black list" );
    const auto& st = *existing;

    require_auth( st.issuer );
    check( sym == st.supply.symbol, "symbol precision mismatch" );
    check( is_account(account), "account does not exist" );

    blacklists blacklist( _self, sym_name );
    auto iter = blacklist.find( account.value );
    if (iter == blacklist.end()) {
        iter = blacklist.emplace( _self, [&]( auto &b ) {
            b.account = account;
        });
    }
}

void token::add_blacklist( const name& addr, const symbol& sym )
{
    blacklists blacklist( get_self(), sym.code().raw() );
    auto iter = blacklist.find( addr.value );
    if( iter == blacklist.end() ) {
        iter = blacklist.emplace( get_self(), [&]( auto &b ) {
            b.account = addr;
        });
    }
}

void token::rmblacklist( const name& account, const symbol& symbol ) {
    auto sym_name = symbol.code().raw();
    stats statstable( _self, sym_name );
    auto existing = statstable.find( sym_name );
    check( existing != statstable.end(), "token with symbol does not exist, create token before remove from black list" );
    const auto& st = *existing;

    require_auth( st.issuer );
    check( symbol == st.supply.symbol, "symbol precision mismatch" );
    blacklists blacklists( _self, sym_name );
    auto it = blacklists.find( account.value );
    check( it != blacklists.end(), "Blacklist row already deleted or never existed. Action won't have any effect." );
    blacklists.erase( it );
}

bool token::in_blacklist(const name& addr, const symbol& sym) {
    blacklists blacklists( _self, sym.code().raw() );
    auto it = blacklists.find( addr.value );
    if (it == blacklists.end()) {
        return false;
    } else {
        return true;
    }
}

void token::addproclist( const name& account, const uint64_t& amount, const symbol& sym ) {
    auto sym_name = sym.code().raw();
    stats statstable( _self, sym_name );
    auto existing = statstable.find( sym_name );
    check( existing != statstable.end(), "token with symbol does not exist, create token before add to black list" );
    const auto& st = *existing;

    require_auth( st.issuer );
    check( sym == st.supply.symbol, "symbol precision mismatch" );
    check( is_account(account), "account does not exist" );

    proclists proclist( _self, sym_name );
    auto iter = proclist.find( account.value );
    if (iter == proclist.end()) {
        iter = proclist.emplace( _self, [&]( auto &b ) {
            b.account = account;
            b.amount = amount;
        });
    }
}

void token::rmproclist( const name& account, const symbol& symbol ) {
    auto sym_name = symbol.code().raw();
    stats statstable( _self, sym_name );
    auto existing = statstable.find( sym_name );
    check( existing != statstable.end(), "token with symbol does not exist, create token before remove from protect list" );
    const auto& st = *existing;

    require_auth( st.issuer );
    check( symbol == st.supply.symbol, "symbol precision mismatch" );
    proclists proclist( _self, sym_name );
    auto it = proclist.find( account.value );
    check( it != proclist.end(), "Protectlist row already deleted or never existed. Action won't have any effect." );
    proclist.erase( it );
}

bool token::in_proclist(const name& addr, const uint64_t& amount, const symbol& sym) {
    proclists proclist( _self, sym.code().raw() );
    auto iter = proclist.find( addr.value );
    if (iter == proclist.end()) {
        return false;
    } else {
        if (amount > iter->amount * 10000) {
            return true;
        } 
        return false;
    }
}

} /// namespace eosio

