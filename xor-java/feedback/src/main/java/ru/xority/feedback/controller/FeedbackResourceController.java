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

import ru.xority.feedback.controller.dto.CreateFeedbackResourceRequest;
import ru.xority.feedback.controller.dto.CreateFeedbackResourceResponse;
import ru.xority.feedback.controller.dto.GetFeedbackResourceResponse;
import ru.xority.feedback.controller.dto.UpdateFeedbackResourceInfoRequest;
import ru.xority.feedback.entity.FeedbackResourceEntity;
import ru.xority.feedback.entity.FeedbackResourceFilter;
import ru.xority.feedback.service.FeedbackResourceService;
import ru.xority.response.SuccessResponse;
import ru.xority.sage.SageHeader;

/**
 * @author foxleren
 */
@RestController
@RequestMapping("/feedback/resource")
@RequiredArgsConstructor
public class FeedbackResourceController {
    private final FeedbackResourceService feedbackResourceService;

    @GetMapping("/list")
    public ResponseEntity<List<GetFeedbackResourceResponse>> list(@RequestParam Optional<UUID> resourceUuid,
                                                                  @RequestParam Optional<String> name,
                                                                  @RequestParam Optional<UUID> createdByUuid,
                                                                  @RequestParam Optional<Boolean> active) {
        FeedbackResourceFilter filter = new FeedbackResourceFilter(
                resourceUuid,
                name,
                createdByUuid,
                active
        );
        List<GetFeedbackResourceResponse> resources = feedbackResourceService.list(filter)
                .stream()
                .map(GetFeedbackResourceResponse::fromFeedbackResourceEntity)
                .toList();
        return ResponseEntity.ok(resources);
    }

    @PostMapping("/create")
    public ResponseEntity<CreateFeedbackResourceResponse> create(@RequestHeader(SageHeader.ACCOUNT_UUID) UUID accountUuid,
                                                                 @RequestBody @Valid CreateFeedbackResourceRequest request) {
        FeedbackResourceEntity resource = FeedbackResourceEntity.fromCreateFeedbackResourceRequest(accountUuid, request);
        UUID uuid = feedbackResourceService.create(resource);
        return ResponseEntity.ok(new CreateFeedbackResourceResponse(uuid));
    }

    @PutMapping("update-info")
    public ResponseEntity<SuccessResponse> updateInfo(@RequestBody @Valid UpdateFeedbackResourceInfoRequest request) {
        feedbackResourceService.updateInfo(
                request.getResourceUuid(),
                request.getName(),
                request.getDescription()
        );
        return SuccessResponse.create200("Feedback resource info is updated");
    }

    @PutMapping("/set-activate/{resourceUuid}")
    public ResponseEntity<SuccessResponse> deactivate(@PathVariable UUID resourceUuid,
                                                      @RequestParam boolean active) {
        feedbackResourceService.setActive(resourceUuid, active);
        return SuccessResponse.create200("Feedback resource active status is updated");
    }
}
