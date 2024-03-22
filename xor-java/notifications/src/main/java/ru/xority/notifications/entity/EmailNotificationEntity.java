package ru.xority.notifications.entity;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.time.LocalDateTime;
import java.util.Objects;
import java.util.UUID;

import lombok.AllArgsConstructor;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class EmailNotificationEntity {
    public static final String UUID_FIELD = "uuid";
    public static final String SENDER_UUID_FIELD = "sender_uuid";
    public static final String SENDER_EMAIL = "sender_email";
    public static final String RECEIVER_EMAIL = "receiver_email";
    public static final String SUBJECT_FIELD = "subject";
    public static final String BODY_FIELD = "body";
    public static final String CREATED_AT = "created_at";

    private UUID uuid;
    private UUID senderUuid;
    private String senderEmail;
    private String receiverEmail;
    private String subject;
    private String body;
    private LocalDateTime createdAt;

    public static EmailNotificationEntity fromResultSet(ResultSet rs) throws SQLException {
        return new EmailNotificationEntity(
                UUID.fromString(Objects.requireNonNull(rs.getString(UUID_FIELD))),
                UUID.fromString(Objects.requireNonNull(rs.getString(SENDER_UUID_FIELD))),
                Objects.requireNonNull(rs.getString(SENDER_EMAIL)),
                Objects.requireNonNull(rs.getString(RECEIVER_EMAIL)),
                Objects.requireNonNull(rs.getString(SUBJECT_FIELD)),
                Objects.requireNonNull(rs.getString(BODY_FIELD)),
                Objects.requireNonNull(rs.getTimestamp(CREATED_AT)).toLocalDateTime()
        );
    }
}
