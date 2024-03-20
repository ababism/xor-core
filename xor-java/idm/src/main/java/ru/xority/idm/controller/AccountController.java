package ru.xority.idm.controller;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import ru.xority.idm.controller.dto.AccountResponse;
import ru.xority.idm.entity.AccountFilter;
import ru.xority.idm.entity.UpdateAccountInfoEntity;
import ru.xority.idm.mapper.AccountMapper;
import ru.xority.idm.service.AccountService;
import ru.xority.response.SuccessResponse;

/**
 * @author foxleren
 */
@RestController
@RequestMapping("/account")
@RequiredArgsConstructor
public class AccountController {
    private final AccountService accountService;
    private final AccountMapper accountMapper;

    @GetMapping("/list")
    public ResponseEntity<?> list(@RequestParam Optional<UUID> uuid,
                                  @RequestParam Optional<String> email,
                                  @RequestParam Optional<Boolean> active) {
        AccountFilter filter = new AccountFilter(uuid, email, active);
        List<AccountResponse> accounts = accountService
                .list(filter)
                .stream()
                .map(accountMapper::accountEntityToAccountResponse)
                .toList();
        return ResponseEntity.ok(accounts);
    }

    @PutMapping("/deactivate/{uuid}")
    public ResponseEntity<?> deactivate(@PathVariable UUID uuid) {
        accountService.deactivate(uuid);
        return SuccessResponse.create200("Account is deactivated");
    }

    @PutMapping("/activate/{uuid}")
    public ResponseEntity<?> activate(@PathVariable UUID uuid) {
        accountService.activate(uuid);
        return SuccessResponse.create200("Account is activated");
    }

    @PutMapping("/update-password/{uuid}")
    public ResponseEntity<?> updatePassword(@PathVariable UUID uuid, @RequestParam String password) {
        accountService.updatePassword(uuid, password);
        return SuccessResponse.create200("Account password is updated");
    }

    @PutMapping("/update-info/{uuid}")
    public ResponseEntity<?> updateInfo(@PathVariable UUID uuid,
                                        @RequestParam Optional<String> firstName,
                                        @RequestParam Optional<String> lastName,
                                        @RequestParam Optional<String> telegramUsername) {
        UpdateAccountInfoEntity entity = new UpdateAccountInfoEntity(
                firstName,
                lastName,
                telegramUsername
        );
        accountService.updateInfo(uuid, entity);
        return SuccessResponse.create200("Account information is updated");
    }
}