package ru.xority.notifications.controller;

import java.util.List;
import java.util.UUID;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpHeaders;
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

import ru.xority.notifications.controller.dto.CreatePlatformNotificationRequest;
import ru.xority.notifications.controller.dto.CreatePlatformNotificationResponse;
import ru.xority.notifications.controller.dto.GetPlatformNotificationResponse;
import ru.xority.notifications.entity.PlatformNotificationEntity;
import ru.xority.notifications.entity.PlatformNotificationFilter;
import ru.xority.notifications.service.PlatformNotificationService;
import ru.xority.response.SuccessResponse;
import ru.xority.sage.SageHeader;

/**
 * @author foxleren
 */
@RestController
@RequestMapping("/platform-notification")
@RequiredArgsConstructor
public class PlatformNotificationController {
    private final PlatformNotificationService platformNotificationService;

    @GetMapping("/list")
    public ResponseEntity<?> list(
            @RequestHeader HttpHeaders headers,
            @RequestParam UUID recieverUuid) {
        System.err.println(headers);
        PlatformNotificationFilter filter = PlatformNotificationFilter.byReceiverUuid(recieverUuid);
        List<GetPlatformNotificationResponse> notifications = platformNotificationService.list(filter)
                .stream()
                .map(GetPlatformNotificationResponse::fromPlatformNotificationEntity)
                .toList();
        return ResponseEntity.ok(notifications);
    }

    @PostMapping("/create")
    public ResponseEntity<?> create(@RequestHeader(SageHeader.XOR_ACCOUNT_UUID) UUID accountUuid,
                                    @RequestBody @Valid CreatePlatformNotificationRequest request) {
        PlatformNotificationEntity notification = PlatformNotificationEntity.fromCreatePlatformNotificationRequest(accountUuid, request);
        UUID uuid = platformNotificationService.create(notification);
        return ResponseEntity.ok(new CreatePlatformNotificationResponse(uuid));
    }

    @PutMapping("/check/{notificationUuid}")
    public ResponseEntity<?> check(@PathVariable UUID notificationUuid) {
        platformNotificationService.check(notificationUuid);
        return SuccessResponse.create200("Platform notification is checked");
    }
}
