package com.finefine.fabric.service;

import org.junit.Assert;
import org.junit.jupiter.api.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import static org.junit.jupiter.api.Assertions.*;

/**
 * @author finefine at: 2020/7/14 3:30 下午
 */
@SpringBootTest
@RunWith(SpringRunner.class)
class CreditContractServiceTest {
    @Autowired
    private CreditContractService creditContractService;

    @Test
    void register() {
        Assert.assertTrue(creditContractService.register("202007140001"));
        Assert.assertTrue(creditContractService.register("202007140002"));
    }

    @Test
    void querybalance() {
        Assert.assertEquals(0,creditContractService.querybalance("202007140001"));
        Assert.assertEquals(0,creditContractService.querybalance("202007140002"));
    }


    @Test
    void transfer() {
        Assert.assertEquals(200l,creditContractService.querybalance("202007140001"));
        Assert.assertTrue(creditContractService.transfer("202007140001","202007140002",80));
        Assert.assertEquals(120l,creditContractService.querybalance("202007140001"));
        Assert.assertEquals(180l,creditContractService.querybalance("202007140002"));
    }

    @Test
    void consume() {
        Assert.assertEquals(120l,creditContractService.querybalance("202007140001"));
        Assert.assertTrue(creditContractService.consume("202007140001",20));
        Assert.assertEquals(100l,creditContractService.querybalance("202007140001"));
    }

    @Test
    void charge() {
        Assert.assertTrue(creditContractService.charge("202007140001",200l));
        Assert.assertTrue(creditContractService.charge("202007140002",100l));
    }
}