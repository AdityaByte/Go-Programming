package com.api.controller;

import java.util.Map;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.api.model.RequestModel;

@RestController
@RequestMapping("/api")
public class MainController {
    
    @GetMapping
    public ResponseEntity<String> handleGetRequest(){
        return ResponseEntity.ok("Welcome to spring framework");
    }

    @GetMapping("/get")
    public ResponseEntity<Map<String,String>> handleGetRequest1(){
        return new ResponseEntity<>(Map.of("message", "Get request response from spring framework"), HttpStatus.OK);
    }

    @PostMapping("/post")
    public ResponseEntity<?> handlePostRequest(@RequestBody RequestModel requestModel){
        return new ResponseEntity<>(requestModel, HttpStatus.OK);
    }

    @PostMapping("/formdata")
    public ResponseEntity<?> handleFormData(
        @ModelAttribute RequestModel model
    ) {
        return new ResponseEntity<>(model, HttpStatus.OK);
    }
}
