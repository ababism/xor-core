package ru.xority.feedback.controller;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import ru.xority.feedback.controller.dto.CreateFeedbackItemRequest;
import ru.xority.feedback.controller.dto.CreateFeedbackItemResponse;
import ru.xority.feedback.controller.dto.GetFeedbackItemResponse;
import ru.xority.feedback.controller.dto.UpdateFeedbackItemInfoRequest;
import ru.xority.feedback.entity.FeedbackItemEntity;
import ru.xority.feedback.entity.FeedbackItemFilter;
import ru.xority.feedback.service.FeedbackItemService;
import ru.xority.response.SuccessResponse;
import ru.xority.sage.SageHeader;

/**
 * @author foxleren
 */
@RestController
@RequestMapping("/feedback/item")
@RequiredArgsConstructor
public class FeedbackItemController {
    private final FeedbackItemService feedbackItemService;

    @GetMapping("/list")
    public ResponseEntity<List<GetFeedbackItemResponse>> list(@RequestParam Optional<UUID> itemUuid,
                                                              @RequestParam Optional<UUID> resourceUuid,
                                                              @RequestParam Optional<UUID> createdByUuid,
                                                              @RequestParam Optional<Boolean> active) {
        FeedbackItemFilter filter = new FeedbackItemFilter(
                itemUuid,
                resourceUuid,
                createdByUuid,
                active
        );
        List<GetFeedbackItemResponse> items = feedbackItemService.list(filter)
                .stream()
                .map(GetFeedbackItemResponse::fromFeedbackItemEntity)
                .toList();
        return ResponseEntity.ok(items);
    }

    @PostMapping("/create")
    public ResponseEntity<CreateFeedbackItemResponse> create(@RequestHeader(SageHeader.ACCOUNT_UUID) UUID accountUuid,
                                                             @RequestBody @Valid CreateFeedbackItemRequest request) {
        FeedbackItemEntity item = FeedbackItemEntity.fromCreateFeedbackItemRequest(accountUuid, request);
        UUID uuid = feedbackItemService.create(item);
        return ResponseEntity.ok(new CreateFeedbackItemResponse(uuid));
    }

    @PutMapping("/update-info")
    public ResponseEntity<SuccessResponse> updateInfo(@RequestBody @Valid UpdateFeedbackItemInfoRequest request) {
        feedbackItemService.updateInfo(
                request.getItemUuid(),
                request.getText(),
                request.getRating()
        );
        return SuccessResponse.create200("Feedback item info is updated");
    }

    @PutMapping("/set-active/{itemUuid}")
    public ResponseEntity<SuccessResponse> deactivate(@PathVariable UUID itemUuid,
                                                      @RequestParam boolean active) {
        feedbackItemService.setActive(itemUuid, active);
        return SuccessResponse.create200("Feedback item active status is updated");
    }
}
