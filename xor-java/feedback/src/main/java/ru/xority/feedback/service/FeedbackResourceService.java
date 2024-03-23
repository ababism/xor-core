package ru.xority.feedback.service;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import ru.xority.feedback.entity.FeedbackResourceEntity;
import ru.xority.feedback.entity.FeedbackResourceFilter;
import ru.xority.feedback.exception.FeedbackResourceNotFoundException;
import ru.xority.feedback.repository.FeedbackResourceRepository;

/**
 * @author foxleren
 */
@Service
@RequiredArgsConstructor
public class FeedbackResourceService {
    private final FeedbackResourceRepository feedbackResourceRepository;

    public List<FeedbackResourceEntity> list(FeedbackResourceFilter filter) {
        return feedbackResourceRepository.list(filter);
    }

    public UUID create(FeedbackResourceEntity resource) {
        feedbackResourceRepository.create(resource);
        return resource.getUuid();
    }

    public void updateInfo(UUID resourceUuid, String name, String description) {
        FeedbackResourceFilter filter = FeedbackResourceFilter.byResourceUuid(resourceUuid);
        Optional<FeedbackResourceEntity> resourceO = feedbackResourceRepository.get(filter);
        resourceO.orElseThrow(FeedbackResourceNotFoundException::new);
        FeedbackResourceEntity resource = resourceO.get();

        resource.setName(name);
        resource.setDescription(description);
        feedbackResourceRepository.update(resource);
    }

    public void setActive(UUID resourceUuid, boolean active) {
        FeedbackResourceFilter filter = FeedbackResourceFilter.byResourceUuid(resourceUuid);
        Optional<FeedbackResourceEntity> resourceO = feedbackResourceRepository.get(filter);
        resourceO.orElseThrow(FeedbackResourceNotFoundException::new);

        FeedbackResourceEntity resource = resourceO.get();
        resource.setActive(active);
        feedbackResourceRepository.update(resource);
    }
}
