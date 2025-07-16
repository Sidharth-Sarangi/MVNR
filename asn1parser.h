#ifndef ASN1PARSER_H
#define ASN1PARSER_H

#include <stdio.h>

typedef struct asn1p_s asn1p_t;
typedef struct asn1p_module_s asn1p_module_t;

// Parsing flags
enum asn1p_flags {
    A1P_NONE             = 0,
    A1P_DEBUG_LEXER      = 1 << 0,
    A1P_DEBUG_PARSER     = 1 << 1,
    A1P_EXTENDED_VALUES  = 1 << 2,
};

// Main parser interface
asn1p_t *asn1p_parse_file(const char *filename, enum asn1p_flags flags);
asn1p_t *asn1p_parse_buffer(const char *buffer, int size, const char *debug_filename, int initial_lineno, enum asn1p_flags flags);
void     asn1p_delete(asn1p_t *a);

extern int asn1p_lineno;
extern const char *asn1p_parse_debug_filename;

#endif
