package ru.xority.notifications.controller.dto;

import java.time.LocalDateTime;
import java.util.UUID;

import lombok.AllArgsConstructor;
import lombok.Data;

import ru.xority.notifications.entity.EmailNotificationEntity;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class GetEmailNotificationResponse {
    private UUID uuid;
    private UUID senderUuid;
    private String senderEmail;
    private String receiverEmail;
    private String subject;
    private String body;
    private LocalDateTime createdAt;

    public static GetEmailNotificationResponse fromEmailNotificationEntity(EmailNotificationEntity entity) {
        return new GetEmailNotificationResponse(
                entity.getUuid(),
                entity.getSenderUuid(),
                entity.getSenderEmail(),
                entity.getReceiverEmail(),
                entity.getSubject(),
                entity.getBody(),
                entity.getCreatedAt()
        );
    }
}
