package ru.xority.idm.account.dao;

import java.util.UUID;

import ru.xority.idm.account.pojo.Account;

/**
 * @author foxleren
 */
public interface AccountDao {
    Account get(UUID id);

    void create(Account account);

    void updatePassword(UUID id, String passwordHash);
}
