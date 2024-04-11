package ru.xority.idm.repository;

import java.sql.Timestamp;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import java.util.UUID;

import lombok.RequiredArgsConstructor;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import ru.xority.idm.entity.RoleEntity;
import ru.xority.idm.entity.RoleFilter;
import ru.xority.sql.SqlQueryHelper;
import ru.xority.utils.DataFilterX;

/**
 * @author foxleren
 */
@Repository
@RequiredArgsConstructor
public class RolePostgresRepository implements RoleRepository {
    private static final String GET = """
            SELECT uuid, name, created_by_uuid, created_at, active FROM role %s;
            """;
    private static final String CREATE = """
            INSERT INTO role
            (uuid, name, created_by_uuid, created_at, active)
            VALUES
            (?, ?, ?, ?, ?);
            """;
    private static final String UPDATE = """
            UPDATE role SET
            name = ?, active = ?
            WHERE uuid = ?;
            """;
    private static final String CREATE_ACCOUNT_ROLE_CONNECTION = """
            INSERT INTO account_role
            (account_uuid, role_uuid, active)
            VALUES
            (?, ?, ?) ON CONFLICT DO NOTHING;
            """;
    private static final String DELETE_ACCOUNT_ROLE_CONNECTION = """
            DELETE FROM account_role
            WHERE account_uuid = ? AND role_uuid = ?;
            """;

    private final JdbcTemplate jdbcTemplate;

    @Override
    public List<RoleEntity> list(RoleFilter filter) {
        Map<String, Object> args = new HashMap<>();
        filter.getUuid().ifPresent(v -> args.put(RoleEntity.UUID_FIELD, v));
        filter.getName().ifPresent(v -> args.put(RoleEntity.NAME_FIELD, v));

        String whereQuery = SqlQueryHelper.queryWhereAnd(args.keySet().stream().toList());

        return jdbcTemplate.query(
                String.format(GET, whereQuery),
                ps -> SqlQueryHelper.buildPreparedStatement(ps, args.values().stream().toList()),
                (rs, i) -> RoleEntity.fromResultSet(rs)
        );
    }

    @Override
    public Optional<RoleEntity> get(RoleFilter filter) {
        List<RoleEntity> role = this.list(filter);
        return DataFilterX.singleO(role);
    }

    @Override
    public void create(RoleEntity role) {
        jdbcTemplate.update(
                CREATE,
                ps -> {
                    ps.setObject(1, role.getUuid());
                    ps.setString(2, role.getName());
                    ps.setObject(3, role.getCreatedByUuid());
                    ps.setTimestamp(4, Timestamp.valueOf(role.getCreatedAt()));
                    ps.setBoolean(5, role.isActive());
                }
        );
    }

    @Override
    public void update(RoleEntity role) {
        jdbcTemplate.update(
                UPDATE,
                ps -> {
                    ps.setString(1, role.getName());
                    ps.setBoolean(2, role.isActive());

                    ps.setObject(3, role.getUuid());

                }
        );
    }

    @Override
    public void assignRole(UUID accountUuid, UUID roleUuid) {
        jdbcTemplate.update(
                CREATE_ACCOUNT_ROLE_CONNECTION,
                ps -> {
                    ps.setObject(1, accountUuid);
                    ps.setObject(2, roleUuid);
                    ps.setBoolean(3, true);
                }
        );
    }

    @Override
    public void revoke(UUID accountUuid, UUID roleUuid) {
        jdbcTemplate.update(
                DELETE_ACCOUNT_ROLE_CONNECTION,
                ps -> {
                    ps.setObject(1, accountUuid);
                    ps.setObject(2, roleUuid);
                }
        );
    }
}
