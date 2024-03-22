package ru.xority.notifications.entity;

import java.util.Optional;
import java.util.UUID;

import lombok.AllArgsConstructor;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class PlatformNotificationFilter {
    private Optional<UUID> notificationUuid;
    private Optional<UUID> senderUuid;
    private Optional<UUID> receiverUuid;

    public static PlatformNotificationFilter byNotificationUuid(UUID notificationUuid) {
        return new PlatformNotificationFilter(
                Optional.of(notificationUuid),
                Optional.empty(),
                Optional.empty()
        );
    }

    public static PlatformNotificationFilter byReceiverUuid(UUID receiverUuid) {
        return new PlatformNotificationFilter(
                Optional.empty(),
                Optional.empty(),
                Optional.of(receiverUuid)
        );
    }
}
