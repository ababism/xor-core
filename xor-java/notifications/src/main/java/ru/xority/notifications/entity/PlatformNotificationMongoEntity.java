package ru.xority.notifications.entity;

import java.time.LocalDateTime;
import java.util.UUID;

import lombok.AllArgsConstructor;
import lombok.Data;
import org.bson.types.ObjectId;
import org.springframework.data.mongodb.core.mapping.MongoId;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class PlatformNotificationMongoEntity {
    public static final String ID_FIELD = "_id";
    public static final String RECEIVER_UUID_FIELD = "receiverUuid";
    public static final String SENDER_UUID_FIELD = "senderUuid";

    @MongoId
    private ObjectId id;
    private UUID receiverUuid;
    private UUID senderUuid;
    private String message;
    private boolean checked;
    private LocalDateTime createdAt;
    private LocalDateTime checkedAt;

    public static PlatformNotificationMongoEntity fromPlatformNotificationEntity(PlatformNotificationEntity platformNotificationEntity) {
        ObjectId id = null;
        if (platformNotificationEntity.getId().isPresent()) {
            id = new ObjectId(platformNotificationEntity.getId().get());
        }
        return new PlatformNotificationMongoEntity(
                id,
                platformNotificationEntity.getReceiverUuid(),
                platformNotificationEntity.getSenderUuid(),
                platformNotificationEntity.getMessage(),
                platformNotificationEntity.isChecked(),
                platformNotificationEntity.getCreatedAt(),
                platformNotificationEntity.getCheckedAt().orElse(null)
        );
    }
}
