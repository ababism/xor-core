package ru.xority.idm.account.dao;

import java.util.List;
import java.util.UUID;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.datatype.jdk8.Jdk8Module;
import org.springframework.dao.support.DataAccessUtils;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import ru.xority.idm.account.pojo.Account;

/**
 * @author foxleren
 */
@Repository
public class AccountPostgresDao implements AccountDao {
    private static final String GET_BY_UUID = """
            SELECT * FROM account WHERE id = ?
            """;

    private static final String CREATE = """
            INSERT INTO account
            (id, login, password_hash, contacts, deleted)
            VALUES(?, ?, ?, ?::jsonb, ?)
            """;
    private static final String UPDATE_PASSWORD = """
            UPDATE account SET password_hash = ? WHERE id = ?
            """;

    private static final ObjectMapper objectMapper = new ObjectMapper().registerModule(new Jdk8Module());

    private final JdbcTemplate jdbcTemplate;

    public AccountPostgresDao(JdbcTemplate jdbcTemplate) {
        this.jdbcTemplate = jdbcTemplate;
    }

    @Override
    public Account get(UUID id) {
        List<Account> accounts = jdbcTemplate.query(
                GET_BY_UUID,
                ps -> {
                    ps.setString(1, id.toString());
                },
                (rs, rowNum) -> Account.fromResultSet(rs)
        );
        return DataAccessUtils.singleResult(accounts);
    }

    @Override
    public void create(Account account) {
        jdbcTemplate.update(
                CREATE,
                ps -> {
                    ps.setObject(1, account.id());
                    ps.setString(2, account.login());
                    ps.setString(3, account.passwordHash());
                    ps.setString(4, account.contacts().toJson(objectMapper));
                    ps.setBoolean(5, account.deleted());
                }
        );
    }

    @Override
    public void updatePassword(UUID id, String passwordHash) {
        jdbcTemplate.update(
                UPDATE_PASSWORD,
                ps -> {
                    ps.setString(1, passwordHash);
                    ps.setObject(2, id);
                }
        );
    }
}
