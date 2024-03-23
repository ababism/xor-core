package ru.xority.feedback.repository;

import java.util.List;
import java.util.Optional;

import lombok.RequiredArgsConstructor;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import ru.xority.feedback.entity.FeedbackResourceEntity;
import ru.xority.feedback.entity.FeedbackResourceFilter;

/**
 * @author foxleren
 */
// TODO implement
@Repository
@RequiredArgsConstructor
public class FeedbackResourcePostgresRepository implements FeedbackResourceRepository {
    private final JdbcTemplate jdbcTemplate;

    @Override
    public List<FeedbackResourceEntity> list(FeedbackResourceFilter filter) {
        return null;
    }

    @Override
    public Optional<FeedbackResourceEntity> get(FeedbackResourceFilter filter) {
        return Optional.empty();
    }

    @Override
    public void create(FeedbackResourceEntity resource) {

    }

    @Override
    public void update(FeedbackResourceEntity resource) {

    }
}
