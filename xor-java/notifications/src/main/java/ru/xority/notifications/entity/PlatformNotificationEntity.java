package ru.xority.notifications.entity;

import java.time.LocalDateTime;
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
    private Optional<String> id;
    private UUID receiverUuid;
    private UUID senderUuid;
    private String message;
    private boolean checked;
    private LocalDateTime createdAt;
    private Optional<LocalDateTime> checkedAt;

    public static PlatformNotificationEntity fromCreatePlatformNotificationRequest(CreatePlatformNotificationRequest request) {
        return new PlatformNotificationEntity(
                Optional.empty(),
                request.getReceiverUuid(),
                request.getSenderUuid(),
                request.getMessage(),
                false,
                LocalDateTime.now(),
                Optional.empty()
        );
    }

    public static PlatformNotificationEntity fromPlatformNotificationMongoEntity(PlatformNotificationMongoEntity platformNotificationMongo) {
        return new PlatformNotificationEntity(
                Optional.of(platformNotificationMongo.getId().toHexString()),
                platformNotificationMongo.getReceiverUuid(),
                platformNotificationMongo.getSenderUuid(),
                platformNotificationMongo.getMessage(),
                platformNotificationMongo.isChecked(),
                platformNotificationMongo.getCreatedAt(),
                Optional.ofNullable(platformNotificationMongo.getCheckedAt())
        );
    }
}
