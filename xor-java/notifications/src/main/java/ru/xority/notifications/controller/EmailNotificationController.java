package ru.xority.notifications.controller;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import ru.xority.notifications.controller.dto.CreateEmailNotificationRequest;
import ru.xority.notifications.controller.dto.CreateEmailNotificationResponse;
import ru.xority.notifications.controller.dto.GetEmailNotificationResponse;
import ru.xority.notifications.entity.EmailNotificationFilter;
import ru.xority.notifications.service.EmailNotificationService;
import ru.xority.sage.SageHeader;

/**
 * @author foxleren
 */
@RestController
@RequestMapping("/email-notification")
@RequiredArgsConstructor
public class EmailNotificationController {
    private final EmailNotificationService emailNotificationService;

    @GetMapping("/list")
    public ResponseEntity<?> list(@RequestParam Optional<UUID> notificationUuid,
                                  @RequestParam Optional<UUID> senderUuid,
                                  @RequestParam Optional<String> senderEmail,
                                  @RequestParam Optional<String> receiverEmail) {
        EmailNotificationFilter filter = new EmailNotificationFilter(
                notificationUuid,
                senderUuid,
                senderEmail,
                receiverEmail
        );
        List<GetEmailNotificationResponse> notifications = emailNotificationService.list(filter)
                .stream()
                .map(GetEmailNotificationResponse::fromEmailNotificationEntity)
                .toList();
        return ResponseEntity.ok(notifications);
    }

    @PostMapping("/create")
    public ResponseEntity<?> create(@RequestHeader(SageHeader.ACCOUNT_UUID) UUID accountUuid,
                                    @RequestBody @Valid CreateEmailNotificationRequest request) {
        UUID uuid = emailNotificationService.create(
                accountUuid,
                request.getReceiverEmail(),
                request.getSubject(),
                request.getBody()
        );
        return ResponseEntity.ok(new CreateEmailNotificationResponse(uuid));
    }
}
