package ru.xority.idm.account.api;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import ru.xority.idm.account.api.dto.RegisterAccountRequest;
import ru.xority.idm.account.api.dto.UpdatePasswordRequest;
import ru.xority.idm.account.manager.AccountManager;

/**
 * @author foxleren
 */
@RestController
@RequestMapping("/account")
public class AccountController {
    private final AccountManager accountManager;

    public AccountController(AccountManager accountManager) {
        this.accountManager = accountManager;
    }

    @PostMapping("/register")
    public ResponseEntity<?> register(@RequestBody RegisterAccountRequest registerAccountRequest) {
        accountManager.register(registerAccountRequest.login(), registerAccountRequest.password());
        return ResponseEntity.ok("Account is registered");
    }

    @PutMapping("/update-password")
    public ResponseEntity<?> updatePassword(@RequestBody UpdatePasswordRequest updatePasswordRequest) {
        accountManager.updatePassword(updatePasswordRequest.id(), updatePasswordRequest.password());
        return ResponseEntity.ok("Password is updated");
    }
}
