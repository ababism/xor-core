package ru.xority.idm.repository;

import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import java.util.UUID;

import lombok.RequiredArgsConstructor;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import ru.xority.idm.entity.AccountEntity;
import ru.xority.idm.entity.AccountFilter;
import ru.xority.idm.entity.RoleEntity;
import ru.xority.sql.SqlQueryHelper;
import ru.xority.utils.DataFilterX;

/**
 * @author foxleren
 */
@Repository
@RequiredArgsConstructor
public class AccountPostgresRepository implements AccountRepository {
    private static final String GET = """
            SELECT uuid, email, password_hash, active, first_name, last_name, telegram_username FROM account %s;
            """;
    private static final String CREATE = """
            INSERT INTO account
            (uuid, email, password_hash, active, first_name, last_name, telegram_username)
            VALUES
            (?, ?, ?, ?, ?, ?, ?);
            """;
    private static final String UPDATE = """
            UPDATE account SET
            email = ?, password_hash = ?, active = ?, first_name = ?, last_name = ?, telegram_username = ?
            WHERE uuid = ?;
            """;
    private static final String GET_ACCOUNT_ROLES = """
            SELECT role.uuid, role.name, role.created_by_uuid, role.created_at, role.active
            FROM role
                JOIN account_role ON role.uuid = account_role.role_uuid
                    AND role.active = true AND account_role.account_uuid = ? AND account_role.active = true
                JOIN account ON account_role.account_uuid = account.uuid;
            """;

    private final JdbcTemplate jdbcTemplate;

    @Override
    public List<AccountEntity> list(AccountFilter filter) {
        Map<String, Object> args = new HashMap<>();
        filter.getUuid().ifPresent(v -> args.put(AccountEntity.UUID_FIELD, v));
        filter.getEmail().ifPresent(v -> args.put(AccountEntity.EMAIL_FIELD, v));
        filter.getActive().ifPresent(v -> args.put(AccountEntity.ACTIVE_FIELD, v));

        String whereQuery = SqlQueryHelper.queryWhereAnd(args.keySet().stream().toList());

        return jdbcTemplate.query(
                String.format(GET, whereQuery),
                ps -> SqlQueryHelper.buildPreparedStatement(ps, args.values().stream().toList()),
                (rs, i) -> AccountEntity.fromResultSet(rs)
        );
    }

    @Override
    public Optional<AccountEntity> get(AccountFilter filter) {
        List<AccountEntity> account = this.list(filter);
        return DataFilterX.singleO(account);
    }

    @Override
    public void create(AccountEntity account) {
        jdbcTemplate.update(
                CREATE,
                ps -> {
                    ps.setObject(1, account.getUuid());
                    ps.setString(2, account.getEmail());
                    ps.setString(3, account.getPasswordHash());
                    ps.setBoolean(4, account.isActive());
                    ps.setString(5, account.getFirstName().orElse(null));
                    ps.setString(6, account.getLastName().orElse(null));
                    ps.setString(7, account.getTelegramUsername().orElse(null));
                }
        );
    }

    @Override
    public void update(AccountEntity account) {
        jdbcTemplate.update(
                UPDATE,
                ps -> {
                    ps.setString(1, account.getEmail());
                    ps.setString(2, account.getPasswordHash());
                    ps.setBoolean(3, account.isActive());
                    ps.setString(4, account.getFirstName().orElse(null));
                    ps.setString(5, account.getLastName().orElse(null));
                    ps.setString(6, account.getTelegramUsername().orElse(null));

                    ps.setObject(7, account.getUuid());
                }
        );
    }

    @Override
    public List<RoleEntity> getActiveRoles(UUID accountUuid) {
        return jdbcTemplate.query(
                GET_ACCOUNT_ROLES,
                ps -> {
                    ps.setObject(1, accountUuid);
                },
                (rs, i) -> RoleEntity.fromResultSet(rs)
        );
    }
}
