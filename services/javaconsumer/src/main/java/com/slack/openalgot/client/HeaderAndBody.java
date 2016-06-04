package com.slack.openalgot.client;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties({"Header"})
public class HeaderAndBody {
	
	@JsonProperty("Body")
	private String body;

	public String getBody() {
		return body;
	}

	public void setBody(String body) {
		this.body = body;
	}

	@Override
	public String toString() {
		return "HeaderAndBody [body=" + body + "]";
	}

}
