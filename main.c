#include <stdio.h>
#include "asn1parser.h"

int main() {
    asn1p_t *tree = asn1p_parse_file("test.asn", 0);
    if (!tree) {
        printf("Parsing failed.\n");
        return 1;
    }

    printf("Parsing succeeded.\n");

    // Optional: print some module info if available
    // (This depends on your struct definition)
    return 0;
}
