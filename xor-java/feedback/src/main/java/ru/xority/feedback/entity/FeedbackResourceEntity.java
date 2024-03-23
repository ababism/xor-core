package ru.xority.feedback.entity;

import java.time.LocalDateTime;
import java.util.UUID;

import lombok.AllArgsConstructor;
import lombok.Data;

import ru.xority.feedback.controller.dto.CreateFeedbackResourceRequest;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class FeedbackResourceEntity {
    private UUID uuid;
    private String name;
    private String description;
    private UUID createdByUuid;
    private LocalDateTime createdAt;
    private boolean active;

    public static FeedbackResourceEntity fromCreateFeedbackResourceRequest(UUID createdByUuid, CreateFeedbackResourceRequest request) {
        return new FeedbackResourceEntity(
                UUID.randomUUID(),
                request.getName(),
                request.getDescription(),
                createdByUuid,
                LocalDateTime.now(),
                true
        );
    }
}
