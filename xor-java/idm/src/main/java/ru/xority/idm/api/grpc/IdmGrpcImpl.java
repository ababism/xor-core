package ru.xority.idm.api.grpc;

import java.util.Optional;

import io.grpc.stub.StreamObserver;
import lombok.RequiredArgsConstructor;
import net.devh.boot.grpc.server.service.GrpcService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;

import ru.xority.idm.common.jwt.JwtService;
import ru.xority.idm.entity.AccountEntity;
import ru.xority.idm.entity.AccountFilter;
import ru.xority.idm.exception.AccountNotFoundException;
import ru.xority.idm.service.AccountService;
import ru.xority.idmproto.IdmGrpc;
import ru.xority.idmproto.VerifyRequest;
import ru.xority.idmproto.VerifyResponse;

/**
 * @author foxleren
 */
@GrpcService
@RequiredArgsConstructor
public class IdmGrpcImpl extends IdmGrpc.IdmImplBase {
    private static final Logger logger = LoggerFactory.getLogger(IdmGrpcImpl.class);

    private final JwtService jwtService;
    private final AccountService accountService;

    @Override
    public void verify(VerifyRequest request, StreamObserver<VerifyResponse> responseObserver) {
        final String email = jwtService.extractEmail(request.getAccessToken());
        UserDetails user = accountService.userDetailsService().loadUserByUsername(email);
        Optional<AccountEntity> accountO = accountService.get(AccountFilter.activeByEmail(user.getUsername()));
        if (accountO.isEmpty()) {
            throw new AccountNotFoundException();
        }
        AccountEntity account = accountO.get();

        VerifyResponse response = VerifyResponse.newBuilder()
                .setAccountUuid(account.getUuid().toString())
                .setAccountEmail(account.getEmail())
                .addAllRoles(user.getAuthorities().stream().map(GrantedAuthority::getAuthority).toList())
                .build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
        logger.info("Access is verified for account_uuid={}. Send access information", account.getUuid());
    }
}