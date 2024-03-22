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
public class EmailNotificationFilter {
    private Optional<UUID> notificationUuid;
    private Optional<UUID> senderUuid;
    private Optional<String> senderEmail;
    private Optional<String> receiverEmail;
}
