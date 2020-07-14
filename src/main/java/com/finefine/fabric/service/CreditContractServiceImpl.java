package com.finefine.fabric.service;

import io.github.ecsoya.fabric.FabricQueryRequest;
import io.github.ecsoya.fabric.FabricQueryResponse;
import io.github.ecsoya.fabric.FabricRequest;
import io.github.ecsoya.fabric.FabricResponse;
import io.github.ecsoya.fabric.config.FabricContext;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.util.Assert;

/**
 * @author finefine at: 2020/7/14 3:08 下午
 */
@Service
public class CreditContractServiceImpl implements CreditContractService {
    public static final String FUNCTION_REGISTER = "register";
    public static final String FUNCTION_QUERY = "query";
    public static final String FUNCTION_TRANSFER = "transfer";
    public static final String FUNCTION_CONSUME = "consume";
    public static final String FUNCTION_CHARGE = "charge";

    @Autowired
    private FabricContext fabricContext;

    @Override
    public boolean register(String account) {
        Assert.notNull(account, "账户不能为空");
        FabricRequest fabricRequest = new FabricRequest(FUNCTION_REGISTER,account);
        if (fabricContext.execute(fabricRequest).isOk()){
            return true;
        }
        return false;
    }

    @Override
    public long querybalance(String account) {
        Assert.notNull(account, "账户不能为空");
        FabricQueryRequest<Long> balanceRequest = new FabricQueryRequest<Long>(Long.class,FUNCTION_QUERY,account);
        FabricQueryResponse<Long> balanceResponse = fabricContext.query(balanceRequest);
        if (balanceResponse.isOk()){
            return balanceResponse.data;
        }
        throw new RuntimeException(balanceResponse.errorMsg);
    }

    @Override
    public boolean transfer(String from, String to, long credit) {
        Assert.notNull(from,"来源地址不能为空");
        Assert.notNull(to,"目标地址不能为空");
        if (credit<=0){
            throw new RuntimeException("转账积分必须大于0");
        }
        FabricRequest fabricRequest = new FabricRequest(FUNCTION_TRANSFER,from,to,String.valueOf(credit));
        FabricResponse fabricResponse = fabricContext.execute(fabricRequest);
        if (fabricResponse.isOk()){
            return true;
        }
        throw new RuntimeException(fabricResponse.errorMsg);
    }

    @Override
    public boolean consume(String account, long credit) {
        Assert.notNull(account,"账户地址不能为空");
        if (credit<=0){
            throw new RuntimeException("消费积分不能小于0");
        }
        FabricRequest fabricRequest = new FabricRequest(FUNCTION_CONSUME,account,String.valueOf(credit));
        FabricResponse fabricResponse = fabricContext.execute(fabricRequest);
        if (fabricResponse.isOk()){
            return true;
        }
        throw new RuntimeException(fabricResponse.errorMsg);
    }

    @Override
    public boolean charge(String account, long credit) {
        Assert.notNull(account,"账户地址不能为空");
        if (credit<=0){
            throw new RuntimeException("充值积分不能小于0");
        }
        FabricRequest fabricRequest = new FabricRequest(FUNCTION_CHARGE,account,String.valueOf(credit));
        FabricResponse fabricResponse = fabricContext.execute(fabricRequest);
        if (fabricResponse.isOk()){
            return true;
        }
        throw new RuntimeException(fabricResponse.errorMsg);
    }
}
