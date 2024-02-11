package ru.xority.idm.account.pojo;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.UUID;

/**
 * @author foxleren
 */
public record Account(
        UUID id,
        String login,
        String passwordHash,
        Contacts contacts,
        boolean deleted)
{

    public static Account getRegisterAccount(String login, String passwordHash) {
        return new Account(
                UUID.randomUUID(),
                login,
                passwordHash,
                Contacts.getEmpty(),
                false
        );
    }

    public static Account fromResultSet(ResultSet rs) throws SQLException {
        return new Account(
                UUID.fromString(rs.getString("id")),
                rs.getString("login"),
                rs.getString("password_hash"),
                Contacts.fromResultSet(rs),
                rs.getBoolean("deleted")
        );
    }
}
