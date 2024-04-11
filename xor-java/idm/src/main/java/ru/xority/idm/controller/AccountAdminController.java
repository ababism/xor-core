package ru.xority.idm.controller;

import java.util.Optional;

import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.annotation.AuthenticationPrincipal;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import ru.xority.idm.common.jwt.JwtService;
import ru.xority.idm.controller.dto.AccountAccessInformation;
import ru.xority.idm.entity.AccountEntity;
import ru.xority.idm.entity.AccountFilter;
import ru.xority.idm.exception.AccountNotFoundException;
import ru.xority.idm.service.AccountService;

/**
 * @author foxleren
 */
@RestController
@RequestMapping("/admin/account")
@RequiredArgsConstructor
public class AccountAdminController {
    private final JwtService jwtService;
    private final AccountService accountService;

    @GetMapping("/verify")
    public ResponseEntity<?> access(@AuthenticationPrincipal UserDetails user,
                                    @RequestHeader("Authorization") String authHeader) {
        System.out.println(user.getAuthorities());
//        String jwtToken = authHeader.substring(7);
//        String email = jwtService.extractEmail(jwtToken);


        Optional<AccountEntity> accountO = accountService.get(AccountFilter.activeByEmail(user.getUsername()));
        if (accountO.isEmpty()) {
            throw new AccountNotFoundException();
        }
        AccountEntity account = accountO.get();
        AccountAccessInformation accountAccessInformation = new AccountAccessInformation(
                account.getUuid(),
                account.getEmail(),
                user.getAuthorities().stream().map(GrantedAuthority::getAuthority).toList()
        );
        System.err.println("Access is verified. Send access information for email=" + account.getEmail());
        return new ResponseEntity<>(accountAccessInformation, HttpStatus.OK);
    }
}
