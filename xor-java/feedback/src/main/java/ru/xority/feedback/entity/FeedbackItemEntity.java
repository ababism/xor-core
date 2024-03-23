package ru.xority.feedback.entity;

import java.util.UUID;

import lombok.AllArgsConstructor;
import lombok.Data;

import ru.xority.feedback.controller.dto.CreateFeedbackItemRequest;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class FeedbackItemEntity {
    private UUID uuid;
    private UUID resourceUuid;
    private UUID createdByUuid;
    private String text;
    private int rating;
    private boolean active;

    public static FeedbackItemEntity fromCreateFeedbackItemRequest(UUID createdByUuid, CreateFeedbackItemRequest request) {
        return new FeedbackItemEntity(
                UUID.randomUUID(),
                request.getResourceUuid(),
                createdByUuid,
                request.getText(),
                request.getRating(),
                true
        );
    }
}
