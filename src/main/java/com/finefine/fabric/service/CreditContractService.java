package com.finefine.fabric.service;

/**
 * @author finefine at: 2020/7/14 3:05 下午
 */
public interface CreditContractService {

    boolean register(String account);

    long querybalance(String account);

    boolean transfer(String from,String to,long credit);

    boolean consume(String account,long credit);

    boolean charge(String account,long credit);
}
