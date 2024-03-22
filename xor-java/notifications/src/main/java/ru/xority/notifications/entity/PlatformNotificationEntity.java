package ru.xority.notifications.entity;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Timestamp;
import java.time.LocalDateTime;
import java.util.Objects;
import java.util.Optional;
import java.util.UUID;

import lombok.AllArgsConstructor;
import lombok.Data;

import ru.xority.notifications.controller.dto.CreatePlatformNotificationRequest;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class PlatformNotificationEntity {
    public static final String UUID_FIELD = "uuid";
    public static final String SENDER_UUID_FIELD = "sender_uuid";
    public static final String RECEIVER_UUID_FIELD = "receiver_uuid";
    public static final String MESSAGE_FIELD = "message";
    public static final String CHECKED_FIELD = "checked";
    public static final String CREATED_AT_FIELD = "created_at";
    public static final String CHECKED_AT_FIELD = "checked_at";

    private UUID uuid;
    private UUID senderUuid;
    private UUID receiverUuid;
    private String message;
    private boolean checked;
    private LocalDateTime createdAt;
    private Optional<LocalDateTime> checkedAt;

    public static PlatformNotificationEntity fromResultSet(ResultSet rs) throws SQLException {
        Timestamp rawCheckedAt = rs.getTimestamp(CHECKED_AT_FIELD);
        Optional<LocalDateTime> checkedAt = Optional.empty();
        if (rawCheckedAt != null) {
            checkedAt = Optional.of(rawCheckedAt.toLocalDateTime());
        }
        return new PlatformNotificationEntity(
                UUID.fromString(Objects.requireNonNull(rs.getString(UUID_FIELD))),
                UUID.fromString(Objects.requireNonNull(rs.getString(SENDER_UUID_FIELD))),
                UUID.fromString(Objects.requireNonNull(rs.getString(RECEIVER_UUID_FIELD))),
                Objects.requireNonNull(rs.getString(MESSAGE_FIELD)),
                rs.getBoolean(CHECKED_FIELD),
                Objects.requireNonNull(rs.getTimestamp(CREATED_AT_FIELD)).toLocalDateTime(),
                checkedAt
        );
    }

    public static PlatformNotificationEntity fromCreatePlatformNotificationRequest(UUID senderUuid, CreatePlatformNotificationRequest request) {
        return new PlatformNotificationEntity(
                UUID.randomUUID(),
                senderUuid,
                request.getReceiverUuid(),
                request.getMessage(),
                false,
                LocalDateTime.now(),
                Optional.empty()
        );
    }
}
