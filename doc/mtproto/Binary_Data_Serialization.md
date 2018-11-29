# Binary Data Serialization
> Forwarded from [Binary Data Serialization](https://core.telegram.org/mtproto/serialize)

MTProto operation requires that elementary and composite data types as well as queries to which such data types are passed as arguments or by which they are returned, be transmitted in binary format (i. e. serialized) .
The TL language is used to describe the data types to be serialized.

## General Definitions
For our purposes, we can identify a type with the set of its (serialized) values understood as strings (finite sequences) of 32-bit numbers (transmitted in little endian order).

Therefore:

- Alphabet (A), in this case, is a set of 32-bit numbers (normally, signed, i. e. between -2^31 and 2^31 - 1).
- Value, in this case, is the same as a string in Alphabet A, i. e. a finite (possibly, empty) sequence of 32-bit numbers. The set of all such sequences is designated as A*.
- Type, for our purposes, is the same as the set of legal values of a type, i. e. some set T which is a subset of A* and is a prefix code (i. e. no element of T may be a prefix for any other element). Therefore, any sequence from A* can contain no more than one prefix that is a member of T.
- Value of Type T is any sequence (value) which is a member of T as a subset of A*.
- Compatible Types are the types T and T’ not intersecting as subsets of A*, such that the union of T and T' is a prefix code.
- Coordinated System of Types is a finite or infinite set of types T_1, …, T_n, …, such that any two types from this set are compatible.
- Data Type is the same as type in the sense of the definition above.
- Functional Type is a type describing a function; it is not a type in the sense of the definition above. Initially, we ignore the existence of functional types and describe only the data types; however, in reality, functional types will later be implemented in some extension of this system using the so-called temporary combinators.

## Combinators, Constructors, Composite Data Types
- Combinator is a function that takes arguments of certain types and returns a value of some other type. We normally look at combinators whose argument and result types are data types (rather than functional types).
Arity (of combinator) is a non-negative integer, the number of combinator arguments.
- Combinator identifier is an identifier beginning with a lowercase Roman letter that uniquely identifies a combinator.
- Combinator number or combinator name is a 32-bit number (i.e., an element of A) that uniquely identifies a combinator. Most often, it is CRC32 of the string containing the combinator description without the final semicolon, and with one space between contiguous lexemes. This always falls in the range from 0x01000000 to 0xffffff00. The highest 256 values are reserved for the so-called temporal-logic combinators used to transmit functions. We frequently denote as combinator the combinator name with single quotes: ‘*combinator*’.
- Combinator description is a string of format combinator_name type_arg_1 ... type_arg_N = type_res; where N stands for the arity of the combinator, type_arg_i is the type of the i-th argument (or rather, a string with the combinator name), and type_res is the combinator value type.
- Constructor is a combinator that cannot be computed (reduced). This is used to represent composite data types. For example, combinator ‘int_tree’ with description int_tree IntTree int IntTree = IntTree, alongside combinator empty_tree = IntTree, may be used to define a composite data type called “IntTree” that takes on values in the form of binary trees with integers as nodes.
- Function (functional combinator) is a combinator which may be computed (reduced) on condition that the requisite number of arguments of requisite types are provided. The result of the computation is an expression consisting of constructors and base type values only.
- Normal form is an expression consisting only of constructors and base type values; that which is normally the result of computing a function.
Type identifier is an identifier that normally starts with a capital letter in Roman script and uniquely identifies the type.
Type number or type name is a 32-bit number that uniquely identifies a type; it normally is the sum of the CRC32 values of the descriptions of the type constructors.
- Description of (composite) Type T is a collection of the descriptions of all constructors that take on Type T values. This is normally written as text with each string containing the description of a single constructor. Here is a description of Type ‘IntTree’, for example:

int_tree IntTree int IntTree = IntTree;
empty_tree = IntTree;

- Polymorphic type is a type whose description contains parameters (*type variables*) in lieu of actual types; approximately, what would be a template in C++. Here is a description of Type List alpha where List is a polymorphic type of arity 1 (i. e., dependent on a single argument), and alpha is a type variable which appears as the constructor’s optional parameter (in curly braces):

cons {alpha:Type} alpha (List alpha) = List alpha;
nil {alpha:Type} = List alpha;

- Value of (composite) Type T is any sequence from A* in the format constr_num arg1 ... argN, where constr_num is the index number of some Constructor C which takes on values of Type T, and arg_i is a value of Type T_i which is the type of the i-th argument to Constructor C. For example, let Combinator int_tree have the index number 17, whereas Combinator empty_tree has the index number 239. Then, the value of Type IntTree is, for example, 17 17 239 1 239 2 239 which is more conveniently written as 'int_tree' 'int_tree' 'empty_tree' 1 'empty_tree' 2 ‘empty_tree’. From the standpoint of a high-level language, this is int_tree (int_tree (empty_tree) 1 (empty_tree)) 2 (empty_tree): IntTree.

- Schema is a collection of all the (composite) data type descriptions. This is used to define some agreed-to system of types.

## Boxed and Bare Types
- Boxed type is a type any value of which starts with the constructor number. Since every constructor has a uniquely determined value type, the first number in any boxed type value uniquely defines its type. This guarantees that the various boxed types in totality make up a coordinated system of types. A boxed type identifier is always capitalized.
- Bare type is a type whose values do not contain a constructor number, which is implied instead. A bare type identifier always coincides with the name of the implied constructor (and therefore, begins with a lowercase letter) which may be padded at the front by the percentage sign (%). In addition, if X is a boxed type with no more than a single constructor, then %X refers to the corresponding bare type. The values of a bare type are identical with the set of number sequences obtained by dropping the first number (i. e., the external constructor index number) from the set of values of the corresponding boxed type (which is the result type of the selected constructor), starting with the selected constructor index number. For example, 3 4 is a value of the int_couple bare type, defined using int_couple int int = IntCouple. The corresponding boxed type is IntCouple; if 404 is the constructor index number for int_couple, then 404 3 4 is the value for the IntCouple boxed type which corresponds to the value of the bare type int_couple (also known as %int_couple and %IntCouple; the latter form is conceptually preferable but longer).

Conceptually, only boxed types should be used everywhere. However, for speed and compactness, bare types have to be used (for instance, an array of 10,000 bare int values is 40,000 bytes long, whereas boxed Int values take up twice as much space; therefore, when transmitting a large array of integer identifiers, say, it is more efficient to use the Vector int type rather than Vector Int). In addition, all base types (int, long, double, string) are bare.

If a boxed type is polymorphic of type arity r, this is also true of any derived bare type. In other words, if one were to define intCouple {alpha:Type} int alpha = IntCouple alpha, then, thereafter, intCouple as an identifier would also be a polymorphic type of arity 1 in combinator (and consequently, in constructor and type) descriptions. The notations intCouple X, %(IntCouple X), and %IntCouple X are equivalent.

## Base Types
Base types exist both as bare (int, long, double, string) and as boxed (Int, Long, Double, String) versions. Their constructor identifiers coincide with the names of the relevant bare types. Their pseudodescriptions have the following appearance:

```
int ? = Int;
long ? = Long;
double ? = Double;
string ? = String;
```

Consequently, the int constructor index number, for example, is the CRC32 of the string "int ? = Int".

The values of bare type int are exactly all the single-element sequences, i. e. numbers between -2^31 and 2^31-1 represent themselves in this case. Values of type long are two-element sequences that are 64-bit signed numbers (little endian again). Values of type double, again, are two-element sequences containing 64-bit real numbers in a standard double format. And finally, the values of type string look differently depending on the length L of the string being serialized:

- If L <= 253, the serialization contains one byte with the value of L, then L bytes of the string followed by 0 to 3 characters containing 0, such that the overall length of the value be divisible by 4, whereupon all of this is interpreted as a sequence of int(L/4)+1 32-bit numbers.
- If L >= 254, the serialization contains byte 254, followed by 3 bytes with the string length L, followed by L bytes of the string, further followed by 0 to 3 null padding bytes.

## Object Pseudotype
The Object pseudotype is a “type” which can take on values that belong to any boxed type in the schema. This helps quickly define such types as list of random items without using polymorphic types. It is best not to abuse this capability since it results in the use of dynamic typing. Nonetheless, it is hard to imagine the data structures that we know from PHP and JSON without using the Object pseudotype.

It is recommended to use TypedObject instead whenever possible:

```
object X:Type value:X = TypedObject;
```

## Built-In Composite Types: Vectors and Associative Arrays
The Vector t polymorphic pseudotype is a “type” whose value is a sequence of values of any type t, either boxed or bare.

```
vector {t:Type} # [ t ] = Vector t;
```

Serialization always uses the same constructor “vector” (const 0x1cb5c415 = crc32("vector t:Type # [ t ] = Vector t”) that is not dependent on the specific value of the variable of type t. The value of the Vector t type is the index number of the relevant constructor number followed by N, the number of elements in the vector, and then by N values of type t. The value of the optional parameter t is not involved in the serialization since it is derived from the result type (always known prior to deserialization).

Polymorphic pseudotypes IntHash t and StrHash t are associative arrays mapping integer and string keys to values of type t. They are, in fact, vectors containing bare 2-tuples (int, t) or (string, t):

```
coupleInt {t:Type} int t = CoupleInt t;
intHash {t:Type} (vector %(CoupleInt t)) = IntHash t;
coupleStr {t:Type} string t = CoupleStr t;
strHash {t:Type} (vector %(CoupleStr t)) = StrHash t;
```

The percentage sign, in this case, means that a bare type that corresponds to the boxed type in parentheses is taken; the boxed type in question must have no more than a single constructor, whatever the values of the parameters.

The keys may be sorted or be in some other order (as in PHP arrays). For associative arrays with sorted keys, the IntSortedHash or StrSortedHash alias is used:

```
intSortedHash {t:Type} (intHash t) = IntSortedHash t;
strSortedHash {t:Type} (strHash t) = StrSortedHash t;
```

## Polymorphic Type Constructors
The constructor of a polymorphic type does not depend on the specific types to which the polymorphic type is applied. When it is computed, optional parameters (normally containing type variables and placed in curly braces) cease to be optional (the curly braces are removed), and, in addition to that, all parenthesis are also removed. Therefore,

```
vector {t:Type} # [ t ] = Vector t;
```

corresponds to the constructor number crc32("vector t:Type # [ t ] = Vector t") = 0x1cb5c415. During (de)serialization, the specific values of the optional variable t are derived from the result type (i. e. the object being serialized or deserialized) that is always known, and are never serialized explicitly.

Previously, it had to be known which specific variable types each polymorphic type will apply to. To accomplish this, the type system used strings of the form

```
polymorphic_type_name type_1 ... type_N;
```

For example,

```
Vector int;
Vector string;
Vector Object;
```

Now they are ignored.

See also polymorphism in TL.

In this case, the Object pseudotype permits using Vector Object to store lists of anything (the values of any boxed types). Since bare types are efficient when short, in practice it is unlikely that cases more complex than the ones cited above will be required.

## Field Names
Let us say that we need to represent users as triplets containing one integer (user ID) and two strings (first and last names). The requisite data structure is the triplet int, string, string which may be declared as follows:

```
user int string string = User;
```

On the other hand, a group may be described by a similar triplet consisting of a group ID, its name, and description:

```
group int string string = Group;
```

For the difference between User and Group to be clear, it is convenient to assign names to some or all of the fields:

```
user id:int first_name:string last_name:string = User;
group id:int title:string description:string = Group;
```

If the User type needs to be extended at a later time by having records with some additional field added to it, it could be accomplished as follows:

```
userv2 id:int unread_messages:int first_name:string last_name:string in_groups:vector int = User;
```

Aside from other things, this approach helps define correct mappings between fields that belong to different constructors of the same type, convert between them as well as convert type values into an associative array with string keys (field names, if defined, are natural choices for such keys).

## TL Language
See TL Language
