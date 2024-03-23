package ru.xority.feedback.repository;

import java.util.List;
import java.util.Optional;

import lombok.RequiredArgsConstructor;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import ru.xority.feedback.entity.FeedbackItemEntity;
import ru.xority.feedback.entity.FeedbackItemFilter;

/**
 * @author foxleren
 */
// TODO implement
@Repository
@RequiredArgsConstructor
public class FeedbackItemPostgresRepository implements FeedbackItemRepository {
    private final JdbcTemplate jdbcTemplate;

    @Override
    public List<FeedbackItemEntity> list(FeedbackItemFilter filter) {
        return null;
    }

    @Override
    public Optional<FeedbackItemEntity> get(FeedbackItemFilter filter) {
        return Optional.empty();
    }

    @Override
    public void create(FeedbackItemEntity item) {

    }

    @Override
    public void update(FeedbackItemEntity item) {

    }
}
