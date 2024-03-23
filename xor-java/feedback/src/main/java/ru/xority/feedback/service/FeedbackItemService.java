package ru.xority.feedback.service;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import ru.xority.exception.BadRequestException;
import ru.xority.feedback.entity.FeedbackItemEntity;
import ru.xority.feedback.entity.FeedbackItemFilter;
import ru.xority.feedback.exception.FeedbackItemNotFoundException;
import ru.xority.feedback.repository.FeedbackItemRepository;

/**
 * @author foxleren
 */
@Service
@RequiredArgsConstructor
public class FeedbackItemService {
    private final FeedbackItemRepository feedbackItemRepository;

    public List<FeedbackItemEntity> list(FeedbackItemFilter filter) {
        return feedbackItemRepository.list(filter);
    }

    public UUID create(FeedbackItemEntity item) {
        FeedbackItemFilter filter = new FeedbackItemFilter(
                Optional.empty(),
                Optional.of(item.getResourceUuid()),
                Optional.of(item.getCreatedByUuid()),
                Optional.empty()
        );
        Optional<FeedbackItemEntity> itemFromRepoO = feedbackItemRepository.get(filter);
        itemFromRepoO.orElseThrow(() -> new BadRequestException("Feedback item is already created"));

        feedbackItemRepository.create(item);
        return item.getUuid();
    }

    public void updateInfo(UUID itemUuid, String text, int rating) {
        FeedbackItemFilter filter = FeedbackItemFilter.byItemUuid(itemUuid);
        Optional<FeedbackItemEntity> itemFromRepoO = feedbackItemRepository.get(filter);
        itemFromRepoO.orElseThrow(FeedbackItemNotFoundException::new);

        FeedbackItemEntity item = itemFromRepoO.get();
        item.setText(text);
        item.setRating(rating);
        feedbackItemRepository.update(item);
    }

    public void deactivate(UUID itemUuid) {
        FeedbackItemFilter filter = FeedbackItemFilter.byItemUuid(itemUuid);
        Optional<FeedbackItemEntity> itemFromRepoO = feedbackItemRepository.get(filter);
        itemFromRepoO.orElseThrow(FeedbackItemNotFoundException::new);

        FeedbackItemEntity item = itemFromRepoO.get();
        if (!item.isActive()) {
            throw new BadRequestException("Feedback item is already deactivated");
        }
        item.setActive(false);
        feedbackItemRepository.update(item);
    }

    public void activate(UUID itemUuid) {
        FeedbackItemFilter filter = FeedbackItemFilter.byItemUuid(itemUuid);
        Optional<FeedbackItemEntity> itemFromRepoO = feedbackItemRepository.get(filter);
        itemFromRepoO.orElseThrow(FeedbackItemNotFoundException::new);

        FeedbackItemEntity item = itemFromRepoO.get();
        if (item.isActive()) {
            throw new BadRequestException("Feedback item is already activated");
        }
        item.setActive(true);
        feedbackItemRepository.update(item);
    }
}
