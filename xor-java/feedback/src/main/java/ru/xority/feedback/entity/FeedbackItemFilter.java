package ru.xority.feedback.entity;

import java.util.Optional;
import java.util.UUID;

import lombok.AllArgsConstructor;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class FeedbackItemFilter {
    private Optional<UUID> itemUuid;
    private Optional<UUID> resourceUuid;
    private Optional<UUID> createdByUuid;
    private Optional<Boolean> active;

    public static FeedbackItemFilter byItemUuid(UUID itemUuid) {
        return new FeedbackItemFilter(
                Optional.of(itemUuid),
                Optional.empty(),
                Optional.empty(),
                Optional.empty()
        );
    }
}
