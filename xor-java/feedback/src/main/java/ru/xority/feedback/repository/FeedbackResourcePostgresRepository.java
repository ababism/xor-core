package ru.xority.feedback.repository;

import java.sql.Timestamp;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Optional;

import lombok.RequiredArgsConstructor;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import ru.xority.feedback.entity.FeedbackResourceEntity;
import ru.xority.feedback.entity.FeedbackResourceFilter;
import ru.xority.sql.SqlQueryHelper;
import ru.xority.utils.DataFilterX;

/**
 * @author foxleren
 */
// TODO implement
@Repository
@RequiredArgsConstructor
public class FeedbackResourcePostgresRepository implements FeedbackResourceRepository {
    private static final String GET = """
            SELECT uuid, name, description, created_by_uuid, created_at, active FROM feedback_resource %s;
            """;
    private static final String CREATE = """
            INSERT INTO feedback_resource
            (uuid, name, description, created_by_uuid, created_at, active)
            VALUES
            (?, ?, ?, ?, ?, ?);
            """;
    private static final String UPDATE = """
            UPDATE feedback_resource SET
            name = ?, description = ?, active = ?
            WHERE uuid = ?;
            """;

    private final JdbcTemplate jdbcTemplate;

    @Override
    public List<FeedbackResourceEntity> list(FeedbackResourceFilter filter) {
        Map<String, Object> args = new HashMap<>();
        filter.getResourceUuid().ifPresent(v -> args.put(FeedbackResourceEntity.UUID_FIELD, v));
        filter.getName().ifPresent(v -> args.put(FeedbackResourceEntity.NAME_FIELD, v));
        filter.getCreatedByUuid().ifPresent(v -> args.put(FeedbackResourceEntity.CREATED_BY_UUID_FIELD, v));
        filter.getActive().ifPresent(v -> args.put(FeedbackResourceEntity.ACTIVE_FIELD, v));

        String whereQuery = SqlQueryHelper.queryWhereAnd(args.keySet().stream().toList());

        return jdbcTemplate.query(
                String.format(GET, whereQuery),
                ps -> SqlQueryHelper.buildPreparedStatement(ps, args.values().stream().toList()),
                (rs, i) -> FeedbackResourceEntity.fromResultSet(rs)
        );
    }

    @Override
    public Optional<FeedbackResourceEntity> get(FeedbackResourceFilter filter) {
        List<FeedbackResourceEntity> resource = this.list(filter);
        return DataFilterX.singleO(resource);
    }

    @Override
    public void create(FeedbackResourceEntity resource) {
        jdbcTemplate.update(
                CREATE,
                ps -> {
                    ps.setObject(1, resource.getUuid());
                    ps.setString(2, resource.getName());
                    ps.setString(3, resource.getDescription());
                    ps.setObject(4, resource.getCreatedByUuid());
                    ps.setTimestamp(5, Timestamp.valueOf(resource.getCreatedAt()));
                    ps.setBoolean(6, resource.isActive());
                }
        );
    }

    @Override
    public void update(FeedbackResourceEntity resource) {
        jdbcTemplate.update(
                UPDATE,
                ps -> {
                    ps.setString(1, resource.getName());
                    ps.setString(2, resource.getDescription());
                    ps.setBoolean(3, resource.isActive());

                    ps.setObject(4, resource.getUuid());
                }
        );
    }
}
