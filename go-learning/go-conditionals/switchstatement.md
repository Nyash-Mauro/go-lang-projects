<div align="center">
A switch statement takes an expression and evaluates it.
For each result of the evaluation, it executes statements. This is known as a case statement.
We don't need break statement here. It is provided automatically by Go language.
default case is executed if any of the case is not satisfied.
</div>


Structure of Switch Statement

+-----------------------+
|  Switch (expression)  | 
+-----------------------+
     |
	 V
	+-------------------+	True	 +-----------------------------+
	| Condition	1		|----------->|	Statements for Condition 1 |----+
	+-------------------+			 +-----------------------------+	|
		| False															|
		| 																|
		V																|
	+-------------------+	True	 +-----------------------------+	|
	| Condition	2		|----------->|	Statements for Condition 2 |----+
	+-------------------+			 +-----------------------------+	|
		| False															|
		| 																|
		V																|
	+-------------------+	True	 +-----------------------------+	|
	| Condition	n		|----------->|	Statements for Condition n |----+
	+-------------------+			 +-----------------------------+	|
		| False															|
		| 																|
		V																|
	+-------------------+	True	 +-----------------------------+	|
	| Default			|----------->|	Default Statements		   |----+
	+-------------------+			 +-----------------------------+	|
		| False															|
		| 																|
		V																|
	+-----------------+													|
	|	  End		  |<------------------------------------------------V
	+-----------------+