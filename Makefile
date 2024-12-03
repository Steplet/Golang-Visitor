antlr = java -jar /home/steplet/vsprojects/Compilers_itmo/antlr-4.10-complete.jar
grammar = Calc.g4

run:
	go run . in.txt

antlr:
	$(antlr) -Dlanguage=Go -visitor -o parser $(grammar)