function input_values() {
  input1 = Number(prompt("Sr. usuario, introduce el primer valor"));
  input2 = Number(prompt("Sr. usuario, introduce el segundo valor"));
  document.getElementById("entrada1").innerHTML = "Primer Valor → " + input1;
  document.getElementById("entrada2").innerHTML = "Segundo Valor → " + input2;
}

function operacion() {
  var x = Number(prompt("Sr. usuario, introduce una opcion"));
  var w = document.getElementById("resultado");
  switch (x) {
    case 0:
      document.body.innerHTML = "¡Se ha salido!";
      break;
    case 1:
      input_values();
      w.innerHTML = input1 + input2;
      break;
    case 2:
      input_values();
      w.innerHTML = input1 - input2;
      break;
    case 3:
      input_values();
      w.innerHTML = input1 * input2;
      break;
  }
}
