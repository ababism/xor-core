package ru.xority.feedback.controller.dto;

import java.time.LocalDateTime;
import java.util.UUID;

import lombok.AllArgsConstructor;
import lombok.Data;

import ru.xority.feedback.entity.FeedbackResourceEntity;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class GetFeedbackResourceResponse {
    private UUID uuid;
    private String name;
    private String description;
    private UUID createdByUuid;
    private LocalDateTime createdAt;
    private boolean active;

    public static GetFeedbackResourceResponse fromFeedbackResourceEntity(FeedbackResourceEntity entity) {
        return new GetFeedbackResourceResponse(
                entity.getUuid(),
                entity.getName(),
                entity.getDescription(),
                entity.getCreatedByUuid(),
                entity.getCreatedAt(),
                entity.isActive()
        );
    }
}
