package ru.xority.idm.entity;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.Optional;
import java.util.UUID;

import lombok.Data;

/**
 * @author foxleren
 */
@Data
public class AccountEntity {
    public static final String UUID_FIELD = "uuid";
    public static final String EMAIL_FIELD = "email";
    public static final String PASSWORD_HASH_FIELD = "password_hash";
    public static final String ACTIVE_FIELD = "active";
    public static final String FIRST_NAME_FIELD = "first_name";
    public static final String LAST_NAME_FIELD = "last_name";
    public static final String TELEGRAM_USERNAME_FIELD = "telegram_username";

    private UUID uuid;
    private String email;
    private String passwordHash;
    private boolean active;
    private Optional<String> firstName;
    private Optional<String> lastName;
    private Optional<String> telegramUsername;

    public static AccountEntity fromResultSet(ResultSet rs) throws SQLException {
        AccountEntity account = new AccountEntity();

        account.setUuid(UUID.fromString(rs.getString(UUID_FIELD)));
        account.setEmail(rs.getString(EMAIL_FIELD));
        account.setPasswordHash(rs.getString(PASSWORD_HASH_FIELD));
        account.setActive(rs.getBoolean(ACTIVE_FIELD));
        account.setFirstName(Optional.ofNullable(rs.getString(FIRST_NAME_FIELD)));
        account.setLastName(Optional.ofNullable(rs.getString(LAST_NAME_FIELD)));
        account.setTelegramUsername(Optional.ofNullable(rs.getString(TELEGRAM_USERNAME_FIELD)));

        return account;
    }

    public static AccountEntity createdAccount(String email, String passwordHash) {
        AccountEntity account = new AccountEntity();

        account.setUuid(UUID.randomUUID());
        account.setEmail(email);
        account.setPasswordHash(passwordHash);
        account.setActive(true);
        account.setFirstName(Optional.empty());
        account.setLastName(Optional.empty());
        account.setTelegramUsername(Optional.empty());

        return account;
    }
}
