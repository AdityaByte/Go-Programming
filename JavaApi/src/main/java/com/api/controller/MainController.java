package com.api.controller;

import java.util.Map;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api")
public class MainController {

    @GetMapping("/health-check")
    public ResponseEntity<String> healthCheck(){
        return ResponseEntity.ok("Health is Good");
    }
    
    @GetMapping
    public ResponseEntity<String> handleGetRequest(){
        return ResponseEntity.ok("Welcome to spring framework");
    }

    @GetMapping("/get")
    public ResponseEntity<Map<String,String>> handleGetRequest1(){
        return new ResponseEntity<>(Map.of("message", "Get request response from spring framework"), HttpStatus.OK);
    }

    @PostMapping("/post")
    public ResponseEntity<?> handlePostRequest(@RequestBody Object o1){
        return new ResponseEntity<>(o1, HttpStatus.OK);
    }

    @PostMapping("/formdata")
    public ResponseEntity<?> handleFormData(
        @RequestParam Map<?,?> formData
    ) {
        return new ResponseEntity<>(formData, HttpStatus.OK);
    }
}
