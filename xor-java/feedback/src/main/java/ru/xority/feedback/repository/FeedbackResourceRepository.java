package ru.xority.feedback.repository;

import java.util.List;
import java.util.Optional;

import ru.xority.feedback.entity.FeedbackResourceEntity;
import ru.xority.feedback.entity.FeedbackResourceFilter;

/**
 * @author foxleren
 */
public interface FeedbackResourceRepository {
    List<FeedbackResourceEntity> list(FeedbackResourceFilter filter);

    Optional<FeedbackResourceEntity> get(FeedbackResourceFilter filter);

    void create(FeedbackResourceEntity resource);

    void update(FeedbackResourceEntity resource);
}
