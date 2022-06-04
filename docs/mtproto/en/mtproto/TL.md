# TL Language
TL (Type Language) serves to describe the used system of types, constructors, and existing functions. In fact, the combinator description format presented in Binary Data Serialization is used.

See also:

Polymorphism in TL
Advanced topics:

Dependent types in TL

Formal description of TL

Formal description of TL combinators

Type serialization

TL schema for serialization of TL schemas

Optional combinator parameters and their values

Binary serialization and abstract TL types

Formal description of templates in TL

Overview
A TL program usually consists of two sections separated by keyword ---functions---. The first section consists of declarations of built-in types and aggregate types (i.e. their constructors). The second section consists of the declared functions, i.e. functional combinators.

Actually, both the first and second sections consist of combinator declarations, each of which ends with a semicolon. However, the first section contains only constructors, while the second section only involves functions. Each combinator is declared using a “combinator declaration” in the format explained above. However, the combinator number and field names may be explicitly assigned.

If additional type declarations are required after functions have been declared, the keyword (section divider) ---types--- is used. Furthermore, a functional combinator may be declared in the type section if its result type begins with an exclamation point (in fact, when the function section is interpreted, this exclamation point is added automatically).

To explicitly define 32-bit names of combinators, a hash mark (#) is added immediately after the combinator's name, followed by 8 hexadecimal digits.

Namespaces
Composite constructions like <namespace_identifier>.<constructor_identifier> and <namespace_identifier>.<Type_identifier> can be used as constructor- or type identifiers. The portion of the identifier to the left of the period is called the namespace. Moreover, the rule about a first uppercase letter in type identifiers and lowercase letter in constructor identifiers applies to the part of the construction after the period. For example, auth.Message would be a type, while auth.std_message would be a constructor.

Namespaces do not require a special declaration.

Comments
Comments are the same as in C++.

Example
// built-in types
int#a8509bda ? = Int;
long ? = Long;
double ? = Double;
string ? = String;
null = Null;

vector {t:Type} # [ t ] = Vector t;
coupleInt {alpha:Type} int alpha = CoupleInt<alpha>;
coupleStr {gamma:Type} string gamma = CoupleStr gamma;  
/* The name of the type variable is irrelevant: "gamma" could be replaced with "alpha";
However, the combinator number will depend on the specific choice. */

intHash {alpha:Type} vector<coupleInt<alpha>> = IntHash<alpha>;
strHash {alpha:Type} (vector (coupleStr alpha)) = StrHash alpha;
intSortedHash {alpha:Type} intHash<alpha> = IntSortedHash<alpha>;
strSortedHash {alpha:Type} (strHash alpha) = StrSortedHash alpha;

// custom types
pair x:Object y:Object = Pair;
triple x:Object y:Object z:Object = Triple;

user#d23c81a3 id:int first_name:string last_name:string = User;
no_user#c67599d1 id:int = User;
group id:int title:string last_name:string = Group;
no_group = Group;

---functions---

// Maybe some built-in arithmetic functions; inverse quotes make "identifiers" out of arbitrary non-alphanumeric strings
`+` Int Int = Int;
`-` Int Int = Int;
`+` Double Double = Double;
// ...

// API functions (aka RPC functions)
getUser#b0f732d5 int = User;
getUsers#2d84d5f5 (Vector int) = Vector User;
In this case, the user constructor has been explicitly assigned a number (0xd23c81a3); In fact, this was not necessary, since this value is the CRC32 of the string "user id:int first_name:string last_name:string = User", which would have been used by default.

Special constructors are not required for Vector int, Vector User, Vector Object, etc. -- the same universal constructor can be used everywhere:

vector#1cb5c415 {t:Type} # [ t ] = Vector t;
Note that when the getUsers (Vector int) = Vector User; constructor number is calculated, the CRC32 of the string "getUsers Vector int = Vector User” is computed (from which all parentheses have been removed).

Notation T0<T1,T2,...,Tn> is syntactic sugar for (T0 (T1) (T2) ... (Tn)). For example, Vector<User> and (Vector User) are entirely interchangeable.

Example of an RPC query
Suppose we want to call getUsers([2,3,4]). This query will be serialized into a sequence of 32-bit integers as follows:

0x2d84d5f5 0x1cb5c415 0x3 0x2 0x3 0x4
Please note that TL serialization yields sequences of 32-bit integers. When it has to be embedded into a byte stream, for example a network packet, each 32-bit integer is represented by four bytes in little-endian order. In this way the above query corresponds to the following byte stream:

F5 D5 84 2D 15 C4 B5 1C 03 00 00 00 02 00 00 00 03 00 00 00 04 00 00 00
The response might look something like this:

0x1cb5c415 0x3 0xd23c81a3 0x2 0x74655005 0x00007265 0x72615006 0x72656b 0xc67599d1 0x3 0xd23c81a3 0x4 0x686f4a04 0x6e 0x656f4403
This roughly corresponds to

[{"id":2,"first_name":"Peter", "last_name":"Parker"},{},{"id":4,"first_name":"John","last_name":"Doe"}]
Note that in both cases the same universal constructor vector#1cb5c415 is used: in the request to serialize the value of type Vector int, and in the serialization of the value of type Vector User in the response. There is no ambiguity because in both cases the type of the value being (de)serialized is known before its (de)serialization begins. For example, after receiving the query, the server sees that the first part is 0x2d84d5f5, which corresponds to the combinator getUsers#2d84d5f5 (Vector int) = Vector User. Thus, it is understood that what follows will be a value of type Vector int. After receiving the response to this query, the client knows that it must receive a value of type Vector User and it deserializes the response accordingly.
