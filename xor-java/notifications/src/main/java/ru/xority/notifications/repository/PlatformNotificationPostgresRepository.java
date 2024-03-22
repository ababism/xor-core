package ru.xority.notifications.repository;

import java.sql.Timestamp;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Optional;

import lombok.RequiredArgsConstructor;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import ru.xority.notifications.entity.PlatformNotificationEntity;
import ru.xority.notifications.entity.PlatformNotificationFilter;
import ru.xority.sql.SqlQueryHelper;
import ru.xority.utils.DataFilterX;

/**
 * @author foxleren
 */
@Repository
@RequiredArgsConstructor
public class PlatformNotificationPostgresRepository implements PlatformNotificationRepository {
    private static final String GET = """
            SELECT uuid, sender_uuid, receiver_uuid, message, checked, created_at, checked_at FROM platform_notification %s;
            """;
    private static final String CREATE = """
            INSERT INTO platform_notification
            (uuid, sender_uuid, receiver_uuid, message, checked, created_at, checked_at)
            VALUES
            (?, ?, ?, ?, ?, ?, ?);
            """;
    private static final String UPDATE = """
            UPDATE platform_notification SET
            sender_uuid = ?, receiver_uuid = ?, message = ?, checked = ?, created_at = ?, checked_at = ?
            WHERE uuid = ?;
            """;

    private final JdbcTemplate jdbcTemplate;

    @Override
    public List<PlatformNotificationEntity> list(PlatformNotificationFilter filter) {
        Map<String, Object> args = new HashMap<>();
        filter.getNotificationUuid().ifPresent(v -> args.put(PlatformNotificationEntity.UUID_FIELD, v));
        filter.getSenderUuid().ifPresent(v -> args.put(PlatformNotificationEntity.SENDER_UUID_FIELD, v));
        filter.getReceiverUuid().ifPresent(v -> args.put(PlatformNotificationEntity.RECEIVER_UUID_FIELD, v));

        String whereQuery = SqlQueryHelper.queryWhereAnd(args.keySet().stream().toList());

        return jdbcTemplate.query(
                String.format(GET, whereQuery),
                ps -> SqlQueryHelper.buildPreparedStatement(ps, args.values().stream().toList()),
                (rs, i) -> PlatformNotificationEntity.fromResultSet(rs)
        );
    }

    @Override
    public Optional<PlatformNotificationEntity> get(PlatformNotificationFilter filter) {
        List<PlatformNotificationEntity> account = this.list(filter);
        return DataFilterX.singleO(account);
    }

    @Override
    public void create(PlatformNotificationEntity notification) {
        jdbcTemplate.update(
                CREATE,
                ps -> {
                    ps.setObject(1, notification.getUuid());
                    ps.setObject(2, notification.getSenderUuid());
                    ps.setObject(3, notification.getReceiverUuid());
                    ps.setString(4, notification.getMessage());
                    ps.setBoolean(5, notification.isChecked());
                    ps.setTimestamp(6, Timestamp.valueOf(notification.getCreatedAt()));
                    ps.setTimestamp(7, notification.getCheckedAt().map(Timestamp::valueOf).orElse(null));
                }
        );
    }

    @Override
    public void update(PlatformNotificationEntity notification) {
        jdbcTemplate.update(
                UPDATE,
                ps -> {
                    ps.setObject(1, notification.getSenderUuid());
                    ps.setObject(2, notification.getReceiverUuid());
                    ps.setString(3, notification.getMessage());
                    ps.setBoolean(4, notification.isChecked());
                    ps.setTimestamp(5, Timestamp.valueOf(notification.getCreatedAt()));
                    ps.setTimestamp(6, notification.getCheckedAt().map(Timestamp::valueOf).orElse(null));

                    ps.setObject(7, notification.getUuid());
                }
        );
    }
}
