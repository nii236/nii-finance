package com.biteagle.tradingalgo;

import static org.junit.Assert.*;

import java.util.Base64;

import org.junit.After;
import org.junit.AfterClass;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;

import com.google.protobuf.InvalidProtocolBufferException;
import com.slack.openalgot.MarketEventProtos.Trade;

public class DeserializationTest {

	@BeforeClass
	public static void setUpBeforeClass() throws Exception {
	}

	@AfterClass
	public static void tearDownAfterClass() throws Exception {
	}

	@Before
	public void setUp() throws Exception {
	}

	@After
	public void tearDown() throws Exception {
	}

	@Test
	public void test() {
		byte[] b = Base64.getDecoder().decode("EQAAAAAAyoBAGWiR7Xw/Nb4/");
		try {
			Trade trade = Trade.parseFrom(b);
			System.out.println(trade.getTime());
			System.out.println(trade.getAmount());
			System.out.println(trade.getPrice());
		} catch (InvalidProtocolBufferException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
			fail();
		}
		
	}

}
