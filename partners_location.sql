PGDMP     
                    z            partners_location    14.2    14.2 
    V           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            W           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            X           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            Y           1262    28212    partners_location    DATABASE     u   CREATE DATABASE partners_location WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_United States.1252';
 !   DROP DATABASE partners_location;
                postgres    false                        3079    28213    postgis 	   EXTENSION     ;   CREATE EXTENSION IF NOT EXISTS postgis WITH SCHEMA public;
    DROP EXTENSION postgis;
                   false            Z           0    0    EXTENSION postgis    COMMENT     ^   COMMENT ON EXTENSION postgis IS 'PostGIS geometry and geography spatial types and functions';
                        false    2            ?            1259    29234    floor_partner_table    TABLE        CREATE TABLE public.floor_partner_table (
    id uuid NOT NULL,
    partner character varying(255),
    rating numeric NOT NULL,
    operating_radius numeric,
    latitude numeric NOT NULL,
    longitude numeric NOT NULL,
    wood boolean,
    carpet boolean,
    tiles boolean,
    CONSTRAINT floor_partner_table_latitude_check CHECK (((latitude >= ('-180'::integer)::numeric) AND (latitude <= (180)::numeric))),
    CONSTRAINT floor_partner_table_longitude_check CHECK (((longitude >= ('-180'::integer)::numeric) AND (longitude <= (180)::numeric))),
    CONSTRAINT floor_partner_table_operating_radius_check CHECK ((operating_radius > (0)::numeric)),
    CONSTRAINT floor_partner_table_rating_check CHECK (((rating >= (1)::numeric) AND (rating <= (5)::numeric)))
);
 '   DROP TABLE public.floor_partner_table;
       public         heap    postgres    false            S          0    29234    floor_partner_table 
   TABLE DATA           ~   COPY public.floor_partner_table (id, partner, rating, operating_radius, latitude, longitude, wood, carpet, tiles) FROM stdin;
    public          postgres    false    215   %       ?          0    28520    spatial_ref_sys 
   TABLE DATA           X   COPY public.spatial_ref_sys (srid, auth_name, auth_srid, srtext, proj4text) FROM stdin;
    public          postgres    false    211   ?       ?           2606    29244 ,   floor_partner_table floor_partner_table_pkey 
   CONSTRAINT     j   ALTER TABLE ONLY public.floor_partner_table
    ADD CONSTRAINT floor_partner_table_pkey PRIMARY KEY (id);
 V   ALTER TABLE ONLY public.floor_partner_table DROP CONSTRAINT floor_partner_table_pkey;
       public            postgres    false    215            S   ?  x?}??n?0???y?H??1K??@?qz(??(?O????i?a?t?A????Y?MA+@^T{1Mq?&?]??0mw??v?????????;8r??r;???]??<	?????砽???@5+5??V%%#t߆?q<o?????tڝ??_?ό7V2???????Y?
%?LH*?d߲-?9g~?}????]S?л?b????{?Ƞ"?*?~?R?ax?N??:?????{z??G?[??d>???"@?3??g???(O?w?ڹ[4???????H?\}{?+?)e???zu?=???s????[?????Q 举Z-??c???????& ??f?????x<????c\}ĵ?c?X????j??,?R??M??2??/?`???C????l6? \??N      ?      x?????? ? ?     