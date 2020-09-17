# Evaluador de expresiones lógicas y aritméticas

Para ejecutar el programa es necesario tener golang instalado, fue elaborado con la versión 15.1

Basta con clonar el repositorio o ejecutar:
> go get github.com/kaarla/yalochallenge

Una vez en el directorio *yalochallenge*, para ver los ejemplos de los evaluadores de expresiones, ejecutar:

> go run logic.go

> go run arithmetic.go

El programa realiza un parseo básico de las expresiones, ejecuta el algoritmo Shunting Yard para pasar de notación infija a postfija. La nueva expresión postfija es procesada con un árbol de sinntaxis abstracta básico para su interpretación.
