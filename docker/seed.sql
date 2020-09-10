--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: vagrant
--

INSERT INTO public.users (id, username, avatar_url, profile_html, profile_markdown) VALUES (1, 'brandony', 'http://avatar.me/brandony', '<h1>Brandon</h1>', '#Brandon');

--
-- Data for Name: boxes; Type: TABLE DATA; Schema: public; Owner: vagrant
--

INSERT INTO public.boxes (id, name, user_id, is_private, created_at, updated_at, short_description, description, description_html, description_markdown, tag, downloads) VALUES (1, 'ubuntu-20.04', 1, true, '2020-09-08 23:03:36.435511', '2020-09-08 23:03:36.435511', 'minimal ubuntu 20.04', 'very minimal ubuntu 20.04', NULL, NULL, 'brandony/ubuntu-20.04', 0);

--
-- Data for Name: versions; Type: TABLE DATA; Schema: public; Owner: vagrant
--

INSERT INTO public.versions (id, version, status, created_at, updated_at, description, description_html, description_markdown, number, release_url, revoke_url, box_id) VALUES (1, '1', 'unreleased', '2020-09-08 23:05:21.542489', '2020-09-08 23:05:21.542489', 'first test', NULL, NULL, '1', NULL, NULL, 1);

--
-- Data for Name: providers; Type: TABLE DATA; Schema: public; Owner: vagrant
--

INSERT INTO public.providers (id, name, hosted, hosted_token, original_url, created_at, updated_at, download_url, version_id) VALUES (1, 'virtualbox', NULL, NULL, NULL, '2020-09-08 23:06:21.501436', '2020-09-08 23:06:21.501436', 'https://downloads.com/mrg/vagrant.box', 1);
