all:
	bison -d asn1p_y.y -o asn1p_y.tab.c
	flex -o lex.yy.c asn1p_l.l
	gcc -o testparser main.c asn1parser.c asn1p_y.tab.c lex.yy.c -I. -Wall

clean:
	rm -f testparser *.tab.* lex.yy.c
