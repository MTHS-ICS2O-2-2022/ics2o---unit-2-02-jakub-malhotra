// Copyright (c) 2020 Mr. Coxall All rights reserved
//
// Created by: Mr. Coxall
// Created on: Sep 2020
// This file contains the JS functions for index.html

/**
 * Find Area and Perimeter
 */
function calculateClicked() {
  document.getElementById("area").innerHTML =
    "<p>The Area is: " + 3 * 5 + " cm²" + "</p>"
  document.getElementById("perimeter").innerHTML =
    "<p>The Perimeter is: " + 2 * (3 + 5) + " cm" + "</p>"
}