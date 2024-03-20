package ru.xority.idm.controller;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import ru.xority.idm.controller.dto.AccountResponse;
import ru.xority.idm.entity.AccountFilter;
import ru.xority.idm.mapper.AccountMapper;
import ru.xority.idm.service.AccountService;

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
        AccountFilter filter = new AccountFilter(uuid, email, active.orElse(true));
        List<AccountResponse> accounts = accountService
                .list(filter)
                .stream()
                .map(accountMapper::accountEntityToAccountResponse)
                .toList();
        return ResponseEntity.ok(accounts);
    }
}
