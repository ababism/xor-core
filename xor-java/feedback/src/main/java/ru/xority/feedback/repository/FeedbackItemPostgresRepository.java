package ru.xority.feedback.repository;

import java.sql.Timestamp;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Optional;

import lombok.RequiredArgsConstructor;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import ru.xority.feedback.entity.FeedbackItemEntity;
import ru.xority.feedback.entity.FeedbackItemFilter;
import ru.xority.sql.SqlQueryHelper;
import ru.xority.utils.DataFilterX;

/**
 * @author foxleren
 */
// TODO implement
@Repository
@RequiredArgsConstructor
public class FeedbackItemPostgresRepository implements FeedbackItemRepository {
    private static final String GET = """
            SELECT uuid, resource_uuid, created_by_uuid, created_at, text, rating, active FROM feedback_item %s;
            """;
    private static final String CREATE = """
            INSERT INTO feedback_item
            (uuid, resource_uuid, created_by_uuid, created_at, text, rating, active)
            VALUES
            (?, ?, ?, ?, ?, ?, ?);
            """;
    private static final String UPDATE = """
            UPDATE feedback_item SET
            text = ?, rating = ?, active = ?
            WHERE uuid = ?;
            """;

    private final JdbcTemplate jdbcTemplate;

    @Override
    public List<FeedbackItemEntity> list(FeedbackItemFilter filter) {
        Map<String, Object> args = new HashMap<>();
        filter.getItemUuid().ifPresent(v -> args.put(FeedbackItemEntity.UUID_FIELD, v));
        filter.getResourceUuid().ifPresent(v -> args.put(FeedbackItemEntity.RESOURCE_UUID_FIELD, v));
        filter.getCreatedByUuid().ifPresent(v -> args.put(FeedbackItemEntity.CREATED_BY_UUID_FIELD, v));
        filter.getActive().ifPresent(v -> args.put(FeedbackItemEntity.ACTIVE_FIELD, v));

        String whereQuery = SqlQueryHelper.queryWhereAnd(args.keySet().stream().toList());

        return jdbcTemplate.query(
                String.format(GET, whereQuery),
                ps -> SqlQueryHelper.buildPreparedStatement(ps, args.values().stream().toList()),
                (rs, i) -> FeedbackItemEntity.fromResultSet(rs)
        );
    }

    @Override
    public Optional<FeedbackItemEntity> get(FeedbackItemFilter filter) {
        List<FeedbackItemEntity> item = this.list(filter);
        return DataFilterX.singleO(item);
    }

    @Override
    public void create(FeedbackItemEntity item) {
        jdbcTemplate.update(
                CREATE,
                ps -> {
                    ps.setObject(1, item.getUuid());
                    ps.setObject(2, item.getResourceUuid());
                    ps.setObject(3, item.getCreatedByUuid());
                    ps.setTimestamp(4, Timestamp.valueOf(item.getCreatedAt()));
                    ps.setString(5, item.getText());
                    ps.setInt(6, item.getRating());
                    ps.setBoolean(7, item.isActive());
                }
        );
    }

    @Override
    public void update(FeedbackItemEntity item) {
        jdbcTemplate.update(
                UPDATE,
                ps -> {
                    ps.setString(1, item.getText());
                    ps.setInt(2, item.getRating());
                    ps.setBoolean(3, item.isActive());
                }
        );
    }
}
