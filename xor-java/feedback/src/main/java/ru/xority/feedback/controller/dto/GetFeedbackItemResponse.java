package ru.xority.feedback.controller.dto;

import java.util.UUID;

import lombok.AllArgsConstructor;
import lombok.Data;

import ru.xority.feedback.entity.FeedbackItemEntity;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class GetFeedbackItemResponse {
    private UUID uuid;
    private UUID resourceUuid;
    private UUID createdByUuid;
    private String text;
    private int rating;
    private boolean active;

    public static GetFeedbackItemResponse fromFeedbackItemEntity(FeedbackItemEntity item) {
        return new GetFeedbackItemResponse(
                item.getUuid(),
                item.getResourceUuid(),
                item.getCreatedByUuid(),
                item.getText(),
                item.getRating(),
                item.isActive()
        );
    }
}
