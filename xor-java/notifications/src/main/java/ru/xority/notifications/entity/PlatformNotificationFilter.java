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
    private Optional<String> id;
    private Optional<UUID> receiverUuid;
    private Optional<UUID> senderUuid;

    public static PlatformNotificationFilter byId(String id) {
        return new PlatformNotificationFilter(
                Optional.of(id),
                Optional.empty(),
                Optional.empty()
        );
    }

    public static PlatformNotificationFilter byReceiverUuid(UUID receiverUuid) {
        return new PlatformNotificationFilter(
                Optional.empty(),
                Optional.of(receiverUuid),
                Optional.empty()
        );
    }
}
