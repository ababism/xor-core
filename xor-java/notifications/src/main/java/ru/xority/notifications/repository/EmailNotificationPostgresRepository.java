package ru.xority.notifications.repository;

import java.sql.Timestamp;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Optional;

import lombok.RequiredArgsConstructor;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import ru.xority.notifications.entity.EmailNotificationEntity;
import ru.xority.notifications.entity.EmailNotificationFilter;
import ru.xority.sql.SqlQueryHelper;
import ru.xority.utils.DataFilterX;

/**
 * @author foxleren
 */
@Repository
@RequiredArgsConstructor
public class EmailNotificationPostgresRepository implements EmailNotificationRepository {
    private static final String GET = """
            SELECT uuid, sender_uuid, sender_email, receiver_email, subject, body, created_at FROM email_notification %s;
            """;
    private static final String CREATE = """
            INSERT INTO email_notification
            (uuid, sender_uuid, sender_email, receiver_email, subject, body, created_at)
            VALUES
            (?, ?, ?, ?, ?, ?, ?);
            """;

    private final JdbcTemplate jdbcTemplate;

    @Override
    public List<EmailNotificationEntity> list(EmailNotificationFilter filter) {
        Map<String, Object> args = new HashMap<>();
        filter.getNotificationUuid().ifPresent(v -> args.put(EmailNotificationEntity.UUID_FIELD, v));
        filter.getSenderUuid().ifPresent(v -> args.put(EmailNotificationEntity.SENDER_UUID_FIELD, v));
        filter.getSenderEmail().ifPresent(v -> args.put(EmailNotificationEntity.SENDER_EMAIL, v));
        filter.getReceiverEmail().ifPresent(v -> args.put(EmailNotificationEntity.RECEIVER_EMAIL, v));

        String whereQuery = SqlQueryHelper.queryWhereAnd(args.keySet().stream().toList());

        return jdbcTemplate.query(
                String.format(GET, whereQuery),
                ps -> SqlQueryHelper.buildPreparedStatement(ps, args.values().stream().toList()),
                (rs, i) -> EmailNotificationEntity.fromResultSet(rs)
        );
    }

    @Override
    public Optional<EmailNotificationEntity> get(EmailNotificationFilter filter) {
        List<EmailNotificationEntity> notification = this.list(filter);
        return DataFilterX.singleO(notification);
    }

    @Override
    public void create(EmailNotificationEntity notification) {
        jdbcTemplate.update(
                CREATE,
                ps -> {
                    ps.setObject(1, notification.getUuid());
                    ps.setObject(2, notification.getSenderUuid());
                    ps.setString(3, notification.getSenderEmail());
                    ps.setString(4, notification.getReceiverEmail());
                    ps.setString(5, notification.getSubject());
                    ps.setString(6, notification.getBody());
                    ps.setTimestamp(7, Timestamp.valueOf(notification.getCreatedAt()));
                }
        );
    }
}
