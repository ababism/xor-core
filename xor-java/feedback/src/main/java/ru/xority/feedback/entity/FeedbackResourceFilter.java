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
public class FeedbackResourceFilter {
    private Optional<UUID> resourceUuid;
    private Optional<String> name;
    private Optional<UUID> createdByUuid;
    private Optional<Boolean> active;

    public static FeedbackResourceFilter byResourceUuid(UUID resourceUuid) {
        return new FeedbackResourceFilter(
                Optional.of(resourceUuid),
                Optional.empty(),
                Optional.empty(),
                Optional.empty()
        );
    }
}