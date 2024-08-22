#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

typedef enum { JSON_NULL, JSON_BOOL, JSON_NUMBER,
    JSON_STRING, JSON_ARRAY, JSON_OBJECT
} JsonType;

typedef struct JsonValue {
    JsonType type;
    union {
        int boolValue;
        double numberValue;
        char *stringValue;
        struct JsonArray *arrayValue;
        struct JsonObject *objectValue;
    };
} JsonValue;

typedef struct JsonArray {
    JsonValue **values;
    size_t size;
} JsonArray;

typedef struct JsonObject {
    char **keys;
    JsonValue **values;
    size_t size;
} JsonObject;

void skipWhitespace(const char **json) {
    while (isspace(**json)) {
        (*json)++;
    }
}

JsonValue *parseValue(const char **json);
char *parseString(const char **json);
double parseNumber(const char **json);

JsonValue *parseJson(const char *json) {
    skipWhitespace(&json);
    return parseValue(&json);
}

JsonValue *parseObject(const char **json) {
    (*json)++;
    JsonObject *object = malloc(sizeof(JsonObject));
    object->keys = NULL;
    object->values = NULL;
    object->size = 0;

    skipWhitespace(json);
    while (**json != '}') {
        skipWhitespace(json);
        char *key = parseString(json);
        skipWhitespace(json);
        if (**json == ':') {
            (*json)++;
            skipWhitespace(json);
            JsonValue *value = parseValue(json);
            object->size++;
            object->keys = realloc(object->keys, object->size * sizeof(char *));
            object->values = realloc(object->values, object->size * sizeof(JsonValue *));
            object->keys[object->size - 1] = key;
            object->values[object->size - 1] = value;
        }
        skipWhitespace(json);
        if (**json == ',') {
            (*json)++;
            skipWhitespace(json);
        }
    }
    (*json)++;

    JsonValue *result = malloc(sizeof(JsonValue));
    result->type = JSON_OBJECT;
    result->objectValue = object;
    return result;
}

JsonValue *parseArray(const char **json) {
    (*json)++;
    JsonArray *array = malloc(sizeof(JsonArray));
    array->values = NULL;
    array->size = 0;

    skipWhitespace(json);
    while (**json != ']') {
        JsonValue *value = parseValue(json);
        array->size++;
        array->values = realloc(array->values, array->size * sizeof(JsonValue *));
        array->values[array->size - 1] = value;
        skipWhitespace(json);
        if (**json == ',') {
            (*json)++;
            skipWhitespace(json);
        }
    }
    (*json)++;

    JsonValue *result = malloc(sizeof(JsonValue));
    result->type = JSON_ARRAY;
    result->arrayValue = array;
    return result;
}

JsonValue *parseValue(const char **json) {
    skipWhitespace(json);
    if (**json == '{') {
        return parseObject(json);
    } else if (**json == '[') {
        return parseArray(json);
    } else if (**json == '"') {
        JsonValue *value = malloc(sizeof(JsonValue));
        value->type = JSON_STRING;
        value->stringValue = parseString(json);
        return value;
    } else if (isdigit(**json) || **json == '-') {
        JsonValue *value = malloc(sizeof(JsonValue));
        value->type = JSON_NUMBER;
        value->numberValue = parseNumber(json);
        return value;
    } else if (strncmp(*json, "true", 4) == 0) {
        *json += 4;
        JsonValue *value = malloc(sizeof(JsonValue));
        value->type = JSON_BOOL;
        value->boolValue = 1;
        return value;
    } else if (strncmp(*json, "false", 5) == 0) {
        *json += 5;
        JsonValue *value = malloc(sizeof(JsonValue));
        value->type = JSON_BOOL;
        value->boolValue = 0;
        return value;
    } else if (strncmp(*json, "null", 4) == 0) {
        *json += 4;
        JsonValue *value = malloc(sizeof(JsonValue));
        value->type = JSON_NULL;
        return value;
    }
    return NULL;
}

char *parseString(const char **json) {
    (*json)++;
    const char *start = *json;
    char *buffer = malloc(strlen(start) + 1);
    char *ptr = buffer;

    while (**json != '"') {
        if (**json == '\\') {
            (*json)++;
            switch (**json) {
                case '\"': *ptr++ = '\"'; break;
                case '\\': *ptr++ = '\\'; break;
                case '/':  *ptr++ = '/';  break;
                case 'b':  *ptr++ = '\b'; break;
                case 'f':  *ptr++ = '\f'; break;
                case 'n':  *ptr++ = '\n'; break;
                case 'r':  *ptr++ = '\r'; break;
                case 't':  *ptr++ = '\t'; break;
                default: break;
            }
        } else {
            *ptr++ = **json;
        }
        (*json)++;
    }
    *ptr = '\0';
    (*json)++;

    return buffer;
}

double parseNumber(const char **json) {
    char *end;
    double number = strtod(*json, &end);
    *json = end;
    return number;
}







void serializeValue(JsonValue *value, char *buffer, size_t *pos);
void serializeArray(JsonArray *array, char *buffer, size_t *pos);
void serializeObject(JsonObject *object, char *buffer, size_t *pos);
void serializeString(const char *string, char *buffer, size_t *pos);

char *serializeJson(JsonValue *value) {
    char *buffer = malloc(1024);
    size_t pos = 0;
    serializeValue(value, buffer, &pos);
    buffer[pos] = '\0';
    return buffer;
}

void serializeValue(JsonValue *value, char *buffer, size_t *pos) {
    switch (value->type) {
        case JSON_NULL:
            strcpy(buffer + *pos, "null");
            *pos += 4;
            break;
        case JSON_BOOL:
            if (value->boolValue) {
                strcpy(buffer + *pos, "true");
                *pos += 4;
            } else {
                strcpy(buffer + *pos, "false");
                *pos += 5;
            }
            break;
        case JSON_NUMBER:
            *pos += sprintf(buffer + *pos, "%f", value->numberValue);
            break;
        case JSON_STRING:
            serializeString(value->stringValue, buffer, pos);
            break;
        case JSON_ARRAY:
            serializeArray(value->arrayValue, buffer, pos);
            break;
        case JSON_OBJECT:
            serializeObject(value->objectValue, buffer, pos);
            break;
    }
}

void serializeArray(JsonArray *array, char *buffer, size_t *pos) {
    buffer[(*pos)++] = '[';
    for (size_t i = 0; i < array->size; i++) {
        serializeValue(array->values[i], buffer, pos);
        if (i < array->size - 1) {
            buffer[(*pos)++] = ',';
        }
    }
    buffer[(*pos)++] = ']';
}

void serializeObject(JsonObject *object, char *buffer, size_t *pos) {
    buffer[(*pos)++] = '{';
    for (size_t i = 0; i < object->size; i++) {
        serializeString(object->keys[i], buffer, pos);
        buffer[(*pos)++] = ':';
        serializeValue(object->values[i], buffer, pos);
        if (i < object->size - 1) {
            buffer[(*pos)++] = ',';
        }
    }
    buffer[(*pos)++] = '}';
}

void serializeString(const char *string, char *buffer, size_t *pos) {
    buffer[(*pos)++] = '"';
    while (*string) {
        if (*string == '"' || *string == '\\' || *string == '/') {
            buffer[(*pos)++] = '\\';
            buffer[(*pos)++] = *string;
        } else if (*string == '\b') {
            buffer[(*pos)++] = '\\';
            buffer[(*pos)++] = 'b';
        } else if (*string == '\f') {
            buffer[(*pos)++] = '\\';
            buffer[(*pos)++] = 'f';
        } else if (*string == '\n') {
            buffer[(*pos)++] = '\\';
            buffer[(*pos)++] = 'n';
        } else if (*string == '\r') {
            buffer[(*pos)++] = '\\';
            buffer[(*pos)++] = 'r';
        } else if (*string == '\t') {
            buffer[(*pos)++] = '\\';
            buffer[(*pos)++] = 't';
        } else {
            buffer[(*pos)++] = *string;
        }
        string++;
    }
    buffer[(*pos)++] = '"';
}

void freeValue(JsonValue *value) {
    if (!value) return;
    switch (value->type) {
        case JSON_STRING:
            free(value->stringValue);
            break;
        case JSON_ARRAY:
            for (size_t i = 0; i < value->arrayValue->size; i++) {
                freeValue(value->arrayValue->values[i]);
            }
            free(value->arrayValue->values);
            free(value->arrayValue);
            break;
        case JSON_OBJECT:
            for (size_t i = 0; i < value->objectValue->size; i++) {
                free(value->objectValue->keys[i]);
                freeValue(value->objectValue->values[i]);
            }
            free(value->objectValue->keys);
            free(value->objectValue->values);
            free(value->objectValue);
            break;
        default:
            break;
    }
    free(value);
}

JsonValue *getObjectValue(JsonObject *object, const char *key) {
    for (size_t i = 0; i < object->size; i++) {
        if (strcmp(object->keys[i], key) == 0) {
            return object->values[i];
        }
    }
    return NULL;
}

void addObjectValue(JsonObject *object, const char *key, JsonValue *value) {
    object->size++;
    object->keys = realloc(object->keys, object->size * sizeof(char *));
    object->values = realloc(object->values, object->size * sizeof(JsonValue *));
    object->keys[object->size - 1] = strdup(key);
    object->values[object->size - 1] = value;
}

void deleteObjectValue(JsonObject *object, const char *key) {
    for (size_t i = 0; i < object->size; i++) {
        if (strcmp(object->keys[i], key) == 0) {
            free(object->keys[i]);
            freeValue(object->values[i]);
            for (size_t j = i; j < object->size - 1; j++) {
                object->keys[j] = object->keys[j + 1];
                object->values[j] = object->values[j + 1];
            }
            object->size--;
            object->keys = realloc(object->keys, object->size * sizeof(char *));
            object->values = realloc(object->values, object->size * sizeof(JsonValue *));
            return;
        }
    }
}

int main() {
    const char *jsonString = "{\n"
                             "    \"person\": {\n"
                             "        \"name\": \"Alice\\\"\",\n"
                             "        \"age\": 32.1,\n"
                             "        \"scores\": [\n"
                             "            85,\n"
                             "            90,\n"
                             "            78.5\n"
                             "        ],\n"
                             "        \"details\": {\n"
                             "            \"isStudent\": false,\n"
                             "            \"isEmployed\": true\n"
                             "        }\n"
                             "    },\n"
                             "\"json\": true,\n"
                             "\"salary\": 19000,\n"
                             "\"phoneNumbers\" : [\"13467543986\",\"13467543986\",\"13467543986\",\"13467543986\",\"13467543986\",\"13467543986\"]\n"
                             "}";


    JsonValue *jsonValue = parseJson(jsonString);

    JsonObject *jsonObject = jsonValue->objectValue;
    JsonValue *personValue = getObjectValue(jsonObject, "person");
    if (personValue && personValue->type == JSON_OBJECT) {
        JsonValue *nameValue = getObjectValue(personValue->objectValue, "name");
        if (nameValue && nameValue->type == JSON_STRING) {
            printf("Name: %s\n", nameValue->stringValue);
        }
    }

    JsonValue *newValue = malloc(sizeof(JsonValue));
    newValue->type = JSON_BOOL;
    newValue->boolValue = 1;
    addObjectValue(jsonObject, "isEmployed", newValue);

    char *serializedJson = serializeJson(jsonValue);
    printf("Serialized JSON: %s\n", serializedJson);

    free(serializedJson);
    freeValue(jsonValue);

    return 0;
}
