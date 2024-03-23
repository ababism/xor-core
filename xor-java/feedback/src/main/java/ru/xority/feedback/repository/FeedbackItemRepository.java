package ru.xority.feedback.repository;

import java.util.List;
import java.util.Optional;

import ru.xority.feedback.entity.FeedbackItemEntity;
import ru.xority.feedback.entity.FeedbackItemFilter;

/**
 * @author foxleren
 */
public interface FeedbackItemRepository {
    List<FeedbackItemEntity> list(FeedbackItemFilter filter);

    Optional<FeedbackItemEntity> get(FeedbackItemFilter filter);

    void create(FeedbackItemEntity item);

    void update(FeedbackItemEntity item);
}
